# Go Web Framework

This project is a web framework built using Go, which includes several key features like migration, websocket, middleware, model, validation, helpers, command execution, error handling, and auto-compilation. It is designed to provide a solid foundation for building web applications with Go.

این پروژه یک فریم‌ورک وب است که با زبان Go نوشته شده و ویژگی‌های مهمی مانند مهاجرت (migration)، وب‌سوکت (websocket)، میدلور (middleware)، مدل (model)، اعتبارسنجی (validation)، کمک‌کننده‌ها (helpers)، اجرای دستورات (command execution)، مدیریت خطا (error handling) و کامپایل خودکار (auto-compilation) را شامل می‌شود. این فریم‌ورک برای ساخت اپلیکیشن‌های وب با Go طراحی شده است.

## 📋 Requirements | پیش‌نیازها

To run this project, you need the following:  
برای اجرای این پروژه، به موارد زیر نیاز دارید:

- Go 1.19 or higher  
- MySQL database for handling data storage  
- Git installed on your system  
- Supported platforms: Linux, macOS, Windows  

### 📦 Install Required Packages | نصب پکیج‌های مورد نیاز

Before running the project, you need to install the required Go packages.  
قبل از اجرای پروژه، باید پکیج‌های مورد نیاز Go را نصب کنید.

Run the following command to install the dependencies:  
برای نصب وابستگی‌ها، دستور زیر را اجرا کنید:

go mod tidy

## 🚀 Installation and Setup | نصب و راه‌اندازی

Follow these steps to set up the project:  
برای راه‌اندازی پروژه مراحل زیر را دنبال کنید:

1. Install the package using `go get`:  (نصب پکیج با استفاده از دستور `go get`)

```bash
go get github.com/sajad-dev/go-web-framework
```
2. Import the package into your project:  (وارد کردن پکیج به پروژه خود)
   
```bash
import "github.com/sajad-dev/go-web-framework"
```

3. Set up the environment variables (e.g., database connection):  (تنظیم متغیرهای محیطی (برای مثال، اتصال به پایگاه داده))

```bash
cp .env.example .env
```

4. Run the project: (اجرای پروژه)

```bash
go run main.go
```
  نمایش پیام راهنما با دستورات موجود

## 🧪 Running Tests | اجرای تست‌ها

To run tests for this project, follow these steps: (برای اجرای تست‌ها در این پروژه، مراحل زیر را دنبال کنید)

1. Run the tests using the `go test` command:  (اجرای تست‌ها با استفاده از دستور `go test`)
```bash
go test ./test/check-system
```
2. To see more detailed output, use the `-v` flag:  (برای مشاهده خروجی دقیق‌تر، از فلگ `-v` استفاده کنید)
 
 ```bash 
   go test -v ./test/check-system
```

## 🧑‍💻 Author | نویسنده

Mohammad Sajad Poorajam (محمد سجاد پورعجم)
