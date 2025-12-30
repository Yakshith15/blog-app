package com.blog.api_gateway.security;

import io.jsonwebtoken.Claims;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.cloud.gateway.support.ServerWebExchangeUtils;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.net.URI;
import java.util.List;
import java.util.Set;

@Configuration
public class JwtAuthenticationFilter {

    private static final List<String> PUBLIC_PATHS = List.of(
            "/api/auth",
            "/actuator/health"
    );

    private final JwtUtil jwtUtil;

    public JwtAuthenticationFilter(JwtUtil jwtUtil) {
        this.jwtUtil = jwtUtil;
    }

    @Bean
    public GlobalFilter jwtFilter() {
        return (exchange, chain) -> {

            if (isPublicPath(exchange)) {
                return chain.filter(exchange);
            }

            String authHeader = exchange.getRequest()
                    .getHeaders()
                    .getFirst(HttpHeaders.AUTHORIZATION);

            if (authHeader == null || !authHeader.startsWith("Bearer ")) {
                return unauthorized(exchange);
            }

            String token = authHeader.substring(7);

            try {
                Claims claims = jwtUtil.validateToken(token);

                ServerWebExchange mutatedExchange = exchange.mutate()
                        .request(r -> r
                                .header("userId",claims.getSubject())
                                .header("emailVerified", claims.get("emailVerified").toString())
                        )
                        .build();

                return chain.filter(mutatedExchange);

            } catch (Exception ex) {
                return unauthorized(exchange);
            }
        };
    }

    /**
     * Check PUBLIC routes using ORIGINAL request path
     * (before gateway rewrites it)
     */
    private boolean isPublicPath(ServerWebExchange exchange) {

        Set<URI> originalUris =
                exchange.getAttribute(
                        ServerWebExchangeUtils.GATEWAY_ORIGINAL_REQUEST_URL_ATTR
                );

        if (originalUris == null || originalUris.isEmpty()) {
            return false;
        }

        String originalPath = originalUris.iterator().next().getPath();

        return originalPath.equals("/api/auth")
                || originalPath.startsWith("/api/auth/")
                || originalPath.equals("/actuator/health");
    }

    private Mono<Void> unauthorized(ServerWebExchange exchange) {
        exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
        return exchange.getResponse().setComplete();
    }
}
