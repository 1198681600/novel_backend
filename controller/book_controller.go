package web

import (
	"net/http"

	"novel_backend/define"
	"novel_backend/service"
)

type (
	IBookController interface {
		CreateBook(w http.ResponseWriter, r *http.Request)
		GetBook(w http.ResponseWriter, r *http.Request)
		UpdateBook(w http.ResponseWriter, r *http.Request)
		DeleteBook(w http.ResponseWriter, r *http.Request)
		ListBooks(w http.ResponseWriter, r *http.Request)
		SearchBooks(w http.ResponseWriter, r *http.Request)
	}

	bookController struct {
		BookService service.IBookService
	}
)

func newBookController(BookService service.IBookService) IBookController {
	return &bookController{
		BookService: BookService,
	}
}

func (c bookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req define.CreateBookRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	book, err := c.BookService.CreateBook(req.Title, req.IsFinish, req.Author, req.Image, req.Introduction, req.Category, req.Metadata)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.CreateBookResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusCreated, define.CreateBookResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusCreated, Message: "Success"},
		Data:         book,
	})
}

func (c bookController) GetBook(w http.ResponseWriter, r *http.Request) {
	var req define.GetBookRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	book, err := c.BookService.GetBook(req.ID)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.GetBookResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.GetBookResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
		Data:         book,
	})
}

func (c bookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var req define.UpdateBookRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	if err := c.BookService.UpdateBook(req.ID, req.Title, req.IsFinish, req.Author); err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.UpdateBookResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.UpdateBookResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
	})
}

func (c bookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var req define.DeleteBookRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	if err := c.BookService.DeleteBook(req.ID); err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.DeleteBookResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.DeleteBookResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
	})
}

func (c bookController) ListBooks(w http.ResponseWriter, r *http.Request) {
	var req define.ListBooksRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	books, err := c.BookService.ListBooks(req.Page, req.PageSize)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.ListBooksResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.ListBooksResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
		Data:         books,
	})
}

func (c bookController) SearchBooks(w http.ResponseWriter, r *http.Request) {
	var req define.SearchBooksRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	books, err := c.BookService.SearchBooks(req.Query)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.SearchBooksResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.SearchBooksResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
		Data:         books,
	})
}
