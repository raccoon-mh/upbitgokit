# upbitApi
> 🚧 개발중인 저장소 입니다. 🚧

> 발생하는 모든 종류의 문제에 대하여 책임지지 않습니다. 주의하며 사용하세요

A Go-based UpBit Kit for interacting with the Upbit Open API. This library simplifies accessing Upbit's features such as trading, retrieving market data, and managing accounts.

# Feture

CLI 및 참조 가능한 go 패키지를 제공합니다. 아래 기능 목록을 참고하세요.

## 주요 기능

### 1. **계정 관리**
- `accountsget`: 전체 계좌 조회

### 2. **API 키 관리**
- `apikeysget`: API 키 리스트 조회

### 3. **캔들 데이터 조회**
- `candlessecondsget [market] [to] [count]`: 초 단위 캔들 데이터 조회
- `candlesminutesget [unit] [market] [to] [count]`: 분 단위 캔들 데이터 조회
- `candlesdaysget [market] [to] [count]`: 일 단위 캔들 데이터 조회
- `candlesweeksget [market] [to] [count]`: 주 단위 캔들 데이터 조회
- `candlesmonthsget [market] [to] [count]`: 월 단위 캔들 데이터 조회

### 4. **마켓 정보**
- `marketallget [is_details]`: 모든 마켓 종목 코드 조회

### 5. **주문 관리**
- `orderschanceget [market]`: 주문 가능 정보 조회
- `orderget [uuid] [identifier]`: 개별 주문 조회
- `orderuuidsget [market] [order_by]`: 주문 리스트 조회
- `orderopenget [market] [page] [limit] [order_by]`: 체결 대기 주문 조회
- `ordersclosedget [market] [state] [start_time] [end_time] [limit] [order_by]`: 종료된 주문 조회
- `ordercanceldelete [uuid] [identifier]`: 주문 취소
- `orderspost [market] [side] [volume] [price] [ord_type] [identifier] [time_in_force]`: 신규 주문 생성

### 6. **입출금 현황**
- `statuswalletget`: 입출금 현황 조회

### 7. **체결 내역 조회**
- `tradesticksget [market] [to] [count] [cursor] [days_ago]`: 최근 체결 내역 조회

### 8. **셸 자동 완성 스크립트 생성**
- `completion [bash|zsh|fish|powershell]`: 셸 자동 완성 스크립트 생성