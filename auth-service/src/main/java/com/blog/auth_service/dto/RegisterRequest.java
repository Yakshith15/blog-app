package com.blog.auth_service.dto;

public class RegisterRequest {

    private String email;
    private String username;
    private String password;

    public String getEmail() {
        return email;
    }

    public String getUsername() {
        return username;
    }

    public String getPassword() {
        return password;
    }
}
