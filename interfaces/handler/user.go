package interfaces

// 修正必須
import (
    "local.package/config"
    "local.package/interfaces/response"
    "local.package/usecase"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
)
// Userに対するHandlerのインターフェース
type UserHandler interface {
    HandleUserGet(http.ResponseWriter, *http.Request, httprouter.Params)
    HandleUserSignup(http.ResponseWriter, *http.Request, httprouter.Params)
}
type userHandler struct {
    userUseCase usecase.UserUseCase
}
//Userデータに関するHandlerを生成
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
    return &userHandler{
        userUseCase: uu,
    }
}
//ユーザ情報取得
func (uh userHandler) HandleUserGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
    // Contextから認証済みのユーザIDを取得
    ctx := request.Context()
    userID := dddcontext.GetUserIDFromContext(ctx)
    //usecaseレイヤを操作して、ユーザデータ取得
    user, err := usecase.UserUsecase{}.SelectByPrimaryKey(config.DB, userID)
    if err != nil {
        response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
        return
    }
    //レスポンスに必要な情報を詰めて返却
    response.JSON(writer, http.StatusOK, user)
}
// ユーザ新規登録
func (uh userHandler) HandleUserSignup(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
    //リクエストボディを取得
    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
        return
    }
    //リクエストボディのパース
    var requestBody userSignupRequest
    json.Unmarshal(body, &requestBody)
    //usecaseの呼び出し
    err = usecase.UserUsecase{}.Insert(config.DB, requestBody.Name, requestBody.Email)
    if err != nil {
        response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
        return
    }
    // レスポンスに必要な情報を詰めて返却
    response.JSON(writer, http.StatusOK, "")
}