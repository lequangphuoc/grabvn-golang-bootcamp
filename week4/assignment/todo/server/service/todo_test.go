package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository/mocks"
	"testing"
)

func TestGetToDo(t *testing.T)  {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.GetTodoRequest{ Id: "123"}
	mockToDoRep.On("Get", req.Id).Return(toDo, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)

	expectedRes := &pb.GetTodoResponse{ Item: toDo}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestCreateTodo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.CreateTodoRequest{ Item: toDo}

	mockToDoRep.On("Insert", req.Item).Return(nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.CreateTodo(nil, req)

	expectedRes := &pb.CreateTodoResponse{ Id: toDo.Id }

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestListToDo(t *testing.T)  {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.ListTodoRequest{}
	mockToDoRep.On("List", req).Return(toDo, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.ListTodo(nil, req)

	expectedRes := &pb.GetTodoResponse{ Item: toDo}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}
