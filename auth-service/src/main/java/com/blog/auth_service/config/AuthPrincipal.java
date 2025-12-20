package com.blog.auth_service.config;

import java.util.UUID;

public class AuthPrincipal {

    private final UUID userId;
    private final boolean emailVerified;

    public AuthPrincipal(UUID userId, boolean emailVerified) {
        this.userId = userId;
        this.emailVerified = emailVerified;
    }

    public UUID getUserId() {
        return userId;
    }

    public boolean isEmailVerified() {
        return emailVerified;
    }
}
