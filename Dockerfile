# syntax=docker/dockerfile:1
FROM alpine
COPY --link nix/store/ynn1by1qdl16q6qwwh2h7zkgrn36c6i8-glibc-2.35-163/ /nix/store/ynn1by1qdl16q6qwwh2h7zkgrn36c6i8-glibc-2.35-163//
ENTRYPOINT [ "/bin/sh" ]
ENV PATH /nix/store/ynn1by1qdl16q6qwwh2h7zkgrn36c6i8-glibc-2.35-163//bin:/bin
ENV PATH /nix/store/ynn1by1qdl16q6qwwh2h7zkgrn36c6i8-glibc-2.35-163//bin:/bin
