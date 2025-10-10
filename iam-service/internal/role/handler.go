package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	roleService Service
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHandler(roleService Service) *Handler {
	return &Handler{
		roleService: roleService,
	}
}

// CreateRole godoc
// @Summary Create a new role
// @Description Create a new role with role name and description
// @Tags roles
// @Accept json
// @Produce json
// @Param request body CreateRoleRequest true "Role data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /roles [post]
func (h *Handler) CreateRole(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	role, err := h.roleService.CreateRole(req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "role already exists" {
			status = http.StatusConflict
		}
		c.JSON(status, ErrorResponse{
			Error:   "creation_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Role created successfully",
		Data:    role,
	})
}

// GetRole godoc
// @Summary Get role by ID
// @Description Get a specific role by its ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /roles/{id} [get]
func (h *Handler) GetRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid role ID format",
		})
		return
	}

	role, err := h.roleService.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "role_not_found",
			Message: "Role not found",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Role retrieved successfully",
		Data:    role,
	})
}

// GetAllRoles godoc
// @Summary Get all roles
// @Description Get a list of all available roles
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /roles [get]
func (h *Handler) GetAllRoles(c *gin.Context) {
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "fetch_failed",
			Message: "Failed to fetch roles",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Roles retrieved successfully",
		Data:    roles,
	})
}

// UpdateRole godoc
// @Summary Update role
// @Description Update an existing role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param request body UpdateRoleRequest true "Role update data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /roles/{id} [put]
func (h *Handler) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid role ID format",
		})
		return
	}

	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	role, err := h.roleService.UpdateRole(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "role not found" {
			status = http.StatusNotFound
		} else if err.Error() == "role already exists" {
			status = http.StatusConflict
		}
		c.JSON(status, ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Role updated successfully",
		Data:    role,
	})
}

// DeleteRole godoc
// @Summary Delete role
// @Description Delete a role by its ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /roles/{id} [delete]
func (h *Handler) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid role ID format",
		})
		return
	}

	err = h.roleService.DeleteRole(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "role not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, ErrorResponse{
			Error:   "delete_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Role deleted successfully",
		Data:    nil,
	})
}
