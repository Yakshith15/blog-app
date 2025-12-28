package com.blog.api_gateway.filter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.cloud.gateway.route.Route;
import org.springframework.cloud.gateway.support.ServerWebExchangeUtils;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.web.server.ServerWebExchange;

import java.util.Optional;
import java.util.UUID;

@Configuration
public class GlobalLoggingFilter {

    private static final Logger log = LoggerFactory.getLogger(GlobalLoggingFilter.class);
    private static final String REQUEST_ID_HEADER = "X-Request-Id";

    @Bean
    public GlobalFilter loggingFilter() {
        return (exchange, chain) -> {

            long startTime = System.currentTimeMillis();

            String requestId = Optional
                    .ofNullable(exchange.getRequest().getHeaders().getFirst(REQUEST_ID_HEADER))
                    .orElse(UUID.randomUUID().toString());

            ServerHttpRequest mutatedRequest = exchange.getRequest()
                    .mutate()
                    .header(REQUEST_ID_HEADER, requestId)
                    .build();

            ServerWebExchange mutatedExchange = exchange.mutate()
                    .request(mutatedRequest)
                    .build();

            Route route = exchange.getAttribute(ServerWebExchangeUtils.GATEWAY_ROUTE_ATTR);
            String routeId = route != null ? route.getId() : "UNKNOWN";

            log.info(
                    "[REQ] id={} method={} path={} route={}",
                    requestId,
                    exchange.getRequest().getMethod(),
                    exchange.getRequest().getURI().getPath(),
                    routeId
            );

            return chain.filter(mutatedExchange)
                    .doOnSuccess(aVoid -> {
                        long duration = System.currentTimeMillis() - startTime;
                        log.info(
                                "[RES] id={} status={} duration={}ms",
                                requestId,
                                exchange.getResponse().getStatusCode(),
                                duration
                        );
                    });
        };
    }
}

