package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"golang_service/models"
	res "golang_service/utils"

	jwt "golang_service/auth"

	"golang_service/contracts"

	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}
func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if govalidator.IsNull(username) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Data can not empty")
		return
	}

	username = models.Santize(username)
	password = models.Santize(password)

	var user models.User
	result := h.DB.Where("user_name = ?", username).First(&user)
	// find username in user table

	if result.Error != nil {
		fmt.Println("err", result.Error)
		res.JSON(w, 400, "Username or Password incorrect")
		return
	}

	// convert interface to string
	hashedPassword := fmt.Sprintf("%v", user.Password)
	fmt.Println("hash", hashedPassword)

	err := models.CheckPasswordHash(hashedPassword, password)

	if err != nil {
		fmt.Println("err", err)
		res.JSON(w, 401, "Username or Password incorrect")
		return
	}

	token, errCreate := jwt.Create(username)

	if errCreate != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}

	res.JSON(w, 200, token)
	return
}

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	fmt.Println("----------", username, email, password)

	if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Data can not empty")
		return
	}

	if !govalidator.IsEmail(email) {
		res.JSON(w, 400, "Email is invalid")
		return
	}
	var user models.User
	username = models.Santize(username)
	email = models.Santize(email)
	password = models.Santize(password)
	err := h.DB.
		Where("user_name = ? and email=?", username, email).
		First(&user).Error

	if err.Error() != "record not found" {
		res.JSON(w, 409, "Bad query")
		return
	}
	// Error here

	if user.UserName != "" {
		res.JSON(w, 409, "User exists")
		return
	}

	password, err = models.Hash(password)

	if err != nil {
		res.JSON(w, 500, "Register has failed")
		return
	}

	user.Email = email
	user.Password = password
	user.UserName = username
	h.DB.Save(&user)
	res.JSON(w, 201, "Register Succesfully")
	return
}
 
func (h handler)  CreatePair(w http.ResponseWriter, r *http.Request){
	creatorName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var creator models.User
	errFindContractOwner := h.DB.Where("user_name = ?", creatorName).First(&creator)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find contract owner!")
		return
	}
	
	
	
}

var adminEntranceFee string
var userSubmitFee string

func (h handler) ChooseAdmin(w http.ResponseWriter, r *http.Request) {
	contractOwnerName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var contractOwner models.User
	errFindContractOwner := h.DB.Where("user_name = ?", contractOwnerName).First(&contractOwner)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find contract owner!")
		return
	}
	adminAddr := r.PostFormValue("adminAddress")
	adminEntranceFee = r.PostFormValue("adminEntranceFee")

	if govalidator.IsNull(adminAddr) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(adminEntranceFee) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	var admin models.User
	errFindAdmin := h.DB.Where("address = ?", adminAddr).First(&admin)
	if errFindAdmin.Error != nil {
		res.JSON(w, 400, "Cannot find admin!")
		return
	}

	//create transaction on blockchain
	conn, auth := contracts.GetContract(contractOwner.PrivateKey, 0)
	n := new(big.Int)
	adminEntranceFees, _ := n.SetString(adminEntranceFee, 10)
	hexAddress := common.HexToAddress(adminAddr)
	result, err := conn.ChooseAdmin(auth, hexAddress, adminEntranceFees)

	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "chooseAdmin has failed")
		return
	}
	res.JSON(w, 201, result)
	return

}

func (h handler) CreateExams(w http.ResponseWriter, r *http.Request) {
	adminName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var admin models.User
	errFindContractOwner := h.DB.Where("user_name = ?", adminName).First(&admin)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find admin!")
		return
	}
	reward := r.PostFormValue("reward")
	userSubmitFee = r.PostFormValue("userSubmitFee")
	if govalidator.IsNull(reward) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(userSubmitFee) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	fmt.Println(adminEntranceFee)
	adminEntranceFees, _ := strconv.ParseUint(adminEntranceFee, 10, 64)

	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, adminEntranceFees)
	n := new(big.Int)
	m := new(big.Int)
	rewards, _ := n.SetString(reward, 10)
	userSubmitFees, _ := m.SetString(userSubmitFee, 10)
	result, err := conn.CreateExams(auth, rewards, userSubmitFees)
	fmt.Println("adminEntranceFee", adminEntranceFees)
	fmt.Println("error", err)

	if err != nil {
		res.JSON(w, 500, "createExams has failed")
		return
	}
	//listen create exam event
	event := contracts.ListenCreateExamEvent(common.HexToAddress(admin.Address))
	var exam models.Exam
	exam.ExamCode = int(event.ExamCode.Int64())
	exam.Reward, _ = strconv.Atoi(reward)
	exam.Fee = int(adminEntranceFees)
	h.DB.Save(&exam)
	res.JSON(w, 201, result)
	return
}

func (h handler) SubmitTest(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	errFindContractOwner := h.DB.Where("user_name = ?", userName).First(&user)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find admin!")
		return
	}
	examCode := r.PostFormValue("examCode")
	answer := r.PostFormValue("answer")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(answer) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	userSubmitFees, _ := strconv.ParseUint(userSubmitFee, 10, 64)

	//create transaction on blockchain
	conn, auth := contracts.GetContract(user.PrivateKey, userSubmitFees)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.SubmitTest(auth, examCodes, answer)

	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "submitTest has failed")
		return
	}

	//listnet submit test event
	submitEvent := contracts.ListenSubmitTestEvent(common.HexToAddress(user.Address))
	var userExam models.UserExam
	var submitTime = time.Unix(submitEvent.SubmitTime.Int64(), 0)
	userExam.Answer = answer
	userExam.UserID = user.ID
	userExam.ExamCode = int(submitEvent.ExamCode.Int64())
	userExam.Fee = int(userSubmitFees)
	userExam.SubmitTime = submitTime
	h.DB.Save(&userExam)
	res.JSON(w, 201, result)
	return
}

func (h handler) EndSubmit(w http.ResponseWriter, r *http.Request) {
	adminName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var admin models.User
	errFindContractOwner := h.DB.Where("user_name = ?", adminName).First(&admin)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find admin!")
		return
	}
	examCode := r.PostFormValue("examCode")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, 0)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.EndSubmit(auth, examCodes)

	if err != nil {
		res.JSON(w, 500, "endSubmit has failed")
		return
	}

	res.JSON(w, 201, result)
	return

}

func (h handler) SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	adminName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var admin models.User
	errFindContractOwner := h.DB.Where("user_name = ?", adminName).First(&admin)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find admin!")
		return
	}
	examCode := r.PostFormValue("examCode")
	answer := r.PostFormValue("answer")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(answer) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, 0)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.SubmitAnswer(auth, examCodes, answer)

	if err != nil {
		res.JSON(w, 500, "submitAnswer has failed")
		return
	}

	//listen submit answer event
	submitAnswerEvent := contracts.ListenSubmitAnswerEvent(common.HexToAddress(admin.Address))
	var exam models.Exam
	errFindExam := h.DB.First(&exam, submitAnswerEvent.ExamCode)
	if errFindExam.Error != nil {
		res.JSON(w, 400, "Cannot find exam!")
		return
	}
	exam.Answer = answer
	h.DB.Save(&exam)

	res.JSON(w, 201, result)
	return

}

func (h handler) CaculateWinner(w http.ResponseWriter, r *http.Request) {
	adminName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var admin models.User
	errFindContractOwner := h.DB.Where("user_name = ?", adminName).First(&admin)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find admin!")
		return
	}
	examCode := r.PostFormValue("examCode")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, 0)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.CaculateWinner(auth, examCodes)

	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "caculateWinner has failed")
		return
	}

	res.JSON(w, 201, result)
	return

}

func (h handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	contractOwnerName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var contractOwner models.User
	errFindContractOwner := h.DB.Where("user_name = ?", contractOwnerName).First(&contractOwner)
	if errFindContractOwner.Error != nil {
		res.JSON(w, 400, "Cannot find contract owner!")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetContract(contractOwner.PrivateKey, 0)
	result, err := conn.Withdraw(auth)

	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Withdraw has failed")
		return
	}

	res.JSON(w, 201, result)
	return
}

func (h handler) CheckAccountBalance(w http.ResponseWriter, r *http.Request) {
	accountAddr := r.PostFormValue("accountAddr")
	conn, _ := contracts.GetContract("", 0)
	result, err := conn.GetAccountBalance(&bind.CallOpts{From: common.HexToAddress(accountAddr)})

	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "CheckBalance has failed")
		return
	}

	res.JSON(w, 201, result)
}
func (h handler) CheckContractBalance(w http.ResponseWriter, r *http.Request) {
	conn, _ := contracts.GetContract("", 0)
	result, err := conn.GetAccountBalance(&bind.CallOpts{})

	if err != nil {
		res.JSON(w, 500, "CheckContractBalance has failed")
		return
	}

	res.JSON(w, 201, result)
}
