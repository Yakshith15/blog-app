package com.blog.auth_service.service;

import com.blog.auth_service.config.JwtUtil;
import com.blog.auth_service.dto.AuthResponse;
import com.blog.auth_service.dto.LoginRequest;
import com.blog.auth_service.dto.RegisterRequest;
import com.blog.auth_service.model.User;
import com.blog.auth_service.repository.UserRepository;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
public class AuthService {

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final JwtUtil jwtUtil;

    public AuthService(
            UserRepository userRepository,
            PasswordEncoder passwordEncoder,
            JwtUtil jwtUtil
    ) {
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
        this.jwtUtil = jwtUtil;
    }

    public UUID register(RegisterRequest request) {

        if (userRepository.existsByEmail(request.getEmail())) {
            throw new IllegalArgumentException("Email already exists");
        }

        if (userRepository.existsByUsername(request.getUsername())) {
            throw new IllegalArgumentException("Username already exists");
        }

        String hashedPassword =
                passwordEncoder.encode(request.getPassword());

        User user = new User(
                UUID.randomUUID(),
                request.getEmail(),
                request.getUsername(),
                hashedPassword
        );

        userRepository.save(user);

        return user.getId();
    }

    public AuthResponse login(LoginRequest request) {

        User user = userRepository.findByEmail(request.getEmail())
                .orElseThrow(() ->
                        new IllegalArgumentException("Invalid credentials")
                );

        if (!passwordEncoder.matches(
                request.getPassword(),
                user.getPasswordHash()
        )) {
            throw new IllegalArgumentException("Invalid credentials");
        }

        String token = jwtUtil.generateToken(
                user.getId(),
                user.getEmail()
        );

        return new AuthResponse(token, 3600);
    }

    public UUID validateToken(String token) {
        return jwtUtil.extractUserId(token);
    }
}
