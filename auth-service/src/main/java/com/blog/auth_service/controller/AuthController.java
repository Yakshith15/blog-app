package com.blog.auth_service.controller;

import com.blog.auth_service.dto.AuthResponse;
import com.blog.auth_service.dto.LoginRequest;
import com.blog.auth_service.dto.RegisterRequest;
import com.blog.auth_service.service.AuthService;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Map;
import java.util.UUID;

@RestController
public class AuthController {

    private final AuthService authService;

    public AuthController(AuthService authService) {
        this.authService = authService;
    }

    @PostMapping("/register")
    public ResponseEntity<?> register(
            @Valid @RequestBody RegisterRequest request
    ) {
        UUID userId = authService.register(request);
        return ResponseEntity
                .status(HttpStatus.CREATED)
                .body(Map.of("userId", userId));
    }

    @PostMapping("/login")
    public ResponseEntity<AuthResponse> login(
            @Valid @RequestBody LoginRequest request
    ) {
        return ResponseEntity.ok(authService.login(request));
    }

    @GetMapping("/validate")
    public ResponseEntity<?> validate(
            @RequestHeader("Authorization") String header
    ) {
        String token = header.substring(7);
        UUID userId = authService.validateToken(token);

        return ResponseEntity.ok(
                Map.of("userId", userId)
        );
    }
}
