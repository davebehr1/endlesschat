FROM rust:1.61 AS builder

COPY . .
RUN cargo build --release


FROM debian:buster-slim
COPY --from=builder ./target/release/chat_server .
CMD ["./chat_server"]