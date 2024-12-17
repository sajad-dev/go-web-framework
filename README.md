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

## 📝 Available Commands | دستورات موجود

Here are the available commands for this project:  
دستورات موجود برای این پروژه عبارتند از:

- `migration`: Run database migrations  
  اجرای مهاجرت‌های پایگاه داده

- `websocket`: Start the websocket server  
  شروع سرور وب‌سوکت

- `middleware`: Use custom middlewares for request handling  
  استفاده از میدلورهای سفارشی برای مدیریت درخواست‌ها

- `model`: Define and manage database models  
  تعریف و مدیریت مدل‌های پایگاه داده

- `validation`: Perform input validation  
  انجام اعتبارسنجی ورودی‌ها

- `helpers`: Use utility functions for various tasks  
  استفاده از توابع کمکی برای وظایف مختلف

- `command`: Execute custom commands in the framework  
  اجرای دستورات سفارشی در فریم‌ورک

- `error`: Handle errors and exceptions in the application  
  مدیریت خطاها و استثناها در اپلیکیشن

- `auto-compiler`: Automatically recompile the project when changes are detected  
  کامپایل خودکار پروژه هنگام شناسایی تغییرات

- `help`: Display the help message with available commands  
  نمایش پیام راهنما با دستورات موجود

## 🧪 Running Tests | اجرای تست‌ها

To run tests for this project, follow these steps: (برای اجرای تست‌ها در این پروژه، مراحل زیر را دنبال کنید)

1. Run the tests using the `go test` command:  (اجرای تست‌ها با استفاده از دستور `go test`)
```bash
go test ./...
```
2. To see more detailed output, use the `-v` flag:  (برای مشاهده خروجی دقیق‌تر، از فلگ `-v` استفاده کنید)
 
 ```bash 
   go test -v ./...
```
3. If you want to run tests for a specific package or file, specify the path:  (اگر می‌خواهید تست‌ها را برای یک پکیج یا فایل خاص اجرا کنید، مسیر آن را مشخص کنید)
```bash
go test ./path/to/package
```
## 🧑‍💻 Author | نویسنده

Mohammad Sajad Poorajam (محمد سجاد پورعجم)
