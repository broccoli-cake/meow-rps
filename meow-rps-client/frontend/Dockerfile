# 1. Nginx 이미지 사용
FROM nginx:alpine

# 2. 빌드된 정적 파일을 Nginx 기본 경로로 복사
COPY build /usr/share/nginx/html

# 3. 기본 Nginx 설정 덮어쓰기 (선택)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 4. 80 포트 열기
EXPOSE 80

# 5. Nginx 실행
CMD ["nginx", "-g", "daemon off;"]

