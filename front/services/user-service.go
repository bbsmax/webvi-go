package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/pelletier/go-toml"
	"io"
	"net/http"
	"os"
	"strings"
	"webvi-go/front/dto"
	"webvi-go/front/model/user"
	"webvi-go/front/utils"
)

type UserService struct{}

var (
	//UserDto 객체생성
	userDto   = user.UserDto{}
	returnMsg = utils.ReturnMessage{}
)

func (u *UserService) Create(db *gorm.DB, requestData *dto.UserRequest) (bool, *utils.ReturnMessage) {

	//1. 이메일로 회원이 존재하는지 검사.
	if isUser, err := userDto.Find(db, requestData); !isUser {
		//회원이 존재하지 않은 경우.
		//2. 회원비밀번호를 sha256, sha512로 변환.
		id := uuid.New()
		password := requestData.Password
		requestData.ID = id.String()
		requestData.Password = password
		//3. 회원데이터 저장
		if err := userDto.Create(db, requestData); err != nil {
			return false, returnMsg.ReturnMsg(fmt.Sprintf("server error [%v]", err), http.StatusInternalServerError, "internal server error")
		}

		return true, returnMsg.ReturnMsg(fmt.Sprintf(`"ID" : "%v"`, requestData.ID), http.StatusCreated, "created")

	} else if isUser {
		//회원이 존재하는 경우.
		return isUser, returnMsg.ReturnMsg(fmt.Sprintf("user email [%v] is exist", requestData.Email), http.StatusInternalServerError, "internal server error")
	} else if err != nil {
		//에러가 발생한 경우.
		return isUser, err
	}

	return true, nil
}

//회원의 상세정보
func (u *UserService) Get(db *gorm.DB, ID string) (*dto.UserResponse, *utils.ReturnMessage) {
	return userDto.Get(db, ID)
}

func (u *UserService) Update(db *gorm.DB, requestData *dto.UserRequest, r *http.Request) (*dto.UserResponse, *utils.ReturnMessage) {
	fmt.Println("useerRequest : ", requestData)
	//1. ID와 password로 해당 회원이 맞는지 확인
	if isUser, err := userDto.Find(db, requestData); !isUser {
		return nil, returnMsg.ReturnMsg(err.Message, http.StatusInternalServerError, "internal server error")
	} else if err != nil {
		return nil, returnMsg.ReturnMsg(err.Message, http.StatusInternalServerError, "internal server error")
	}

	//2. 회원이 맞으면 회원정보 업데이트
	//2-1 파일이 있는지 확인. 파일 있으면 업로드

	file, header, err := r.FormFile("avator")

	switch err {
	case nil:
		//업로드 파일이 있는 경우.
		config, err := toml.LoadFile("webvi-go/front/config.toml")
		if err != nil {
			return nil, returnMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
		}

		fmt.Println("file : ", file)
		fileExts := strings.Split(header.Filename, ".")
		fileExt := fileExts[len(fileExts)-1]
		fileName := requestData.ID + "." + fileExt
		requestData.Avator = fileName

		if err != nil {
			return nil, returnMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
		}

		defer file.Close()

		avatorUploadPath := config.Get("file.avator").(string)
		out, err := os.Create(avatorUploadPath + fileName)
		if err != nil {
			return nil, returnMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
		}

		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			return nil, returnMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
		}
	case http.ErrMissingFile:
		//업로드 파일이 없는 경우.
		fmt.Println("no file")
	default:

	}

	updateData := u.updateUser(requestData)
	//fmt.Println("updateData : ", updateData)
	return userDto.Update(db, requestData.ID, updateData)

}

func (u *UserService) Delete(db *gorm.DB, ID string) (string, *utils.ReturnMessage) {
	//1. ID와 password로 해당 회원이 맞는지 확인
	//2. 회원이 맞으면 회원정보 업데이트
	return userDto.Delete(db, ID)
}

func (u *UserService) updateUser(requestData *dto.UserRequest) map[string]string {
	returnData := map[string]string{}

	if requestData.Name != "" {
		returnData["name"] = requestData.Name
	}

	if requestData.Password != "" {
		returnData["password"] = requestData.Password
	}

	if requestData.Phone != "" {
		returnData["phone"] = requestData.Phone
	}

	if requestData.Avator != "" {
		returnData["avator"] = requestData.Avator
	}
	return returnData
}
