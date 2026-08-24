-- =============================================================
-- Mini LMS Platform — Seed Data
-- Database: lms-backend
--
-- Default credentials:
--   Admin     admin@cybersoft.edu.vn  /  Admin@123
--   Students  *@email.com             /  Student@123
--
-- Import:
--   mysql -h 127.0.0.1 -P 3310 -u root -p lms-backend < seed.sql
-- =============================================================

USE `lms-backend`;

-- -------------------------------------------------------
-- 1. Users
-- -------------------------------------------------------
INSERT INTO `users` (`id`, `email`, `password`, `full_name`, `role`, `status`, `created_at`, `updated_at`) VALUES
(1, 'admin@cybersoft.edu.vn',
    '$2a$10$ovmAIWvAnz1xGpSaNXsfpeFXCkseHKJ9OvKT0sNRhBiqJp1j7B8l6',
    'Admin CyberSoft', 'admin', 'active', NOW(), NOW()),
(2, 'an.nguyen@email.com',
    '$2a$10$4FN5HCwLjQqhegQu82yMgejsFp5Jr.jQyU.iZikd8Ke4yiRXcpkmW',
    'Nguyễn Văn An', 'student', 'active', NOW(), NOW()),
(3, 'binh.tran@email.com',
    '$2a$10$4FN5HCwLjQqhegQu82yMgejsFp5Jr.jQyU.iZikd8Ke4yiRXcpkmW',
    'Trần Thị Bình', 'student', 'active', NOW(), NOW()),
(4, 'chi.le@email.com',
    '$2a$10$4FN5HCwLjQqhegQu82yMgejsFp5Jr.jQyU.iZikd8Ke4yiRXcpkmW',
    'Lê Minh Chí', 'student', 'active', NOW(), NOW());

-- -------------------------------------------------------
-- 2. Categories
-- -------------------------------------------------------
INSERT INTO `categories` (`id`, `name`, `slug`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Lập trình Web',    'web',      'HTML, CSS, JavaScript, ReactJS, NodeJS và các công nghệ web hiện đại', NOW(), NOW()),
(2, 'Lập trình Mobile', 'mobile',   'iOS, Android, React Native, Flutter — phát triển ứng dụng di động',   NOW(), NOW()),
(3, 'DevOps & Cloud',   'devops',   'Docker, Kubernetes, CI/CD, AWS, GCP — vận hành và triển khai hệ thống', NOW(), NOW()),
(4, 'Cơ sở dữ liệu',   'database', 'MySQL, PostgreSQL, MongoDB, Redis — thiết kế và tối ưu database',      NOW(), NOW()),
(5, 'Thiết kế UI/UX',  'ui-ux',    'Figma, Adobe XD — thiết kế giao diện và trải nghiệm người dùng',       NOW(), NOW());

-- -------------------------------------------------------
-- 3. Courses  (5 courses, all published, 4 lessons each)
-- -------------------------------------------------------
INSERT INTO `courses` (`id`, `category_id`, `title`, `slug`, `description`, `thumbnail`, `price`, `level`, `status`, `total_lessons`, `created_at`, `updated_at`) VALUES
(1, 1, 'JavaScript ES6+ Toàn tập',
    'javascript-es6-toan-tap',
    'Khóa học JavaScript từ cơ bản đến ES6+. Học viên sẽ nắm vững arrow function, destructuring, spread operator, Promise, async/await và các tính năng hiện đại của JavaScript.',
    'https://cybersoft.edu.vn/wp-content/uploads/2021/01/js.jpg',
    699000, 'beginner', 'published', 4, NOW(), NOW()),

(2, 1, 'ReactJS từ cơ bản đến nâng cao',
    'reactjs-tu-co-ban-den-nang-cao',
    'Xây dựng ứng dụng web hiện đại với ReactJS. Từ Component, Props, State đến Hooks, Redux Toolkit và React Router. Học qua dự án thực tế.',
    'https://cybersoft.edu.vn/wp-content/uploads/2021/01/react.jpg',
    899000, 'intermediate', 'published', 4, NOW(), NOW()),

(3, 1, 'Node.js & Express — Backend từ A đến Z',
    'nodejs-express-backend-tu-a-den-z',
    'Xây dựng REST API backend chuyên nghiệp với Node.js và Express. Kết nối database, xác thực JWT, middleware, upload file và deploy lên cloud.',
    'https://cybersoft.edu.vn/wp-content/uploads/2021/01/nodejs.jpg',
    799000, 'intermediate', 'published', 4, NOW(), NOW()),

(4, 2, 'React Native — Lập trình Mobile Cross-platform',
    'react-native-lap-trinh-mobile-cross-platform',
    'Xây dựng ứng dụng mobile iOS và Android từ một codebase duy nhất bằng React Native. Từ thiết lập môi trường, UI components đến navigation và call API.',
    'https://cybersoft.edu.vn/wp-content/uploads/2021/01/react-native.jpg',
    999000, 'intermediate', 'published', 4, NOW(), NOW()),

(5, 3, 'Docker & Kubernetes cho Developer',
    'docker-kubernetes-cho-developer',
    'Làm chủ containerization với Docker và orchestration với Kubernetes. Xây dựng pipeline CI/CD, deploy ứng dụng lên cloud. Kỹ năng DevOps thiết yếu cho developer hiện đại.',
    'https://cybersoft.edu.vn/wp-content/uploads/2021/01/docker.jpg',
    1199000, 'advanced', 'published', 4, NOW(), NOW());

-- -------------------------------------------------------
-- 4. Sections  (2 per course = 10 total)
-- -------------------------------------------------------
INSERT INTO `sections` (`id`, `course_id`, `title`, `sort_order`, `created_at`, `updated_at`) VALUES
-- Course 1: JavaScript
(1,  1, 'Nền tảng JavaScript',          1, NOW(), NOW()),
(2,  1, 'ES6+ và tính năng hiện đại',   2, NOW(), NOW()),
-- Course 2: ReactJS
(3,  2, 'React Cơ bản',                 1, NOW(), NOW()),
(4,  2, 'Xây dựng ứng dụng thực tế',   2, NOW(), NOW()),
-- Course 3: Node.js
(5,  3, 'Node.js Fundamentals',          1, NOW(), NOW()),
(6,  3, 'Database & Authentication',     2, NOW(), NOW()),
-- Course 4: React Native
(7,  4, 'React Native Cơ bản',          1, NOW(), NOW()),
(8,  4, 'Navigation & API Integration', 2, NOW(), NOW()),
-- Course 5: Docker
(9,  5, 'Docker Fundamentals',           1, NOW(), NOW()),
(10, 5, 'Kubernetes Orchestration',      2, NOW(), NOW());

-- -------------------------------------------------------
-- 5. Lessons  (2 per section = 4 per course = 20 total)
-- -------------------------------------------------------
INSERT INTO `lessons` (`id`, `section_id`, `course_id`, `title`, `content`, `duration_minutes`, `sort_order`, `is_preview`, `created_at`, `updated_at`) VALUES

-- Course 1: JavaScript
(1,  1, 1, 'Giới thiệu về JavaScript và môi trường chạy',
    'JavaScript là ngôn ngữ lập trình phổ biến nhất thế giới. Trong bài học này, bạn sẽ hiểu JavaScript chạy như thế nào trên trình duyệt và Node.js, cách thiết lập môi trường phát triển với VS Code và Node.js.',
    20, 1, 1, NOW(), NOW()),
(2,  1, 1, 'Biến, kiểu dữ liệu và toán tử',
    'Khám phá sự khác biệt giữa var, let và const. Tìm hiểu các kiểu dữ liệu primitive (string, number, boolean, null, undefined, symbol) và reference (object, array, function). Thực hành với các toán tử số học, so sánh và logic.',
    30, 2, 0, NOW(), NOW()),
(3,  2, 1, 'Arrow function, Destructuring và Spread operator',
    'Arrow function giúp viết code ngắn gọn hơn và không binding this. Destructuring cho phép trích xuất giá trị từ array và object một cách thanh lịch. Spread operator (...) dùng để copy và merge array/object.',
    35, 1, 0, NOW(), NOW()),
(4,  2, 1, 'Promise, Async/Await và xử lý bất đồng bộ',
    'JavaScript là single-threaded nhưng xử lý async qua Event Loop. Promise giải quyết callback hell. Async/Await làm cho code async trở nên dễ đọc như code đồng bộ. Học cách handle lỗi với try/catch.',
    45, 2, 0, NOW(), NOW()),

-- Course 2: ReactJS
(5,  3, 2, 'Component, Props và State trong React',
    'React xây dựng UI dựa trên component. Mỗi component có thể nhận Props từ cha và quản lý State nội bộ. Tìm hiểu sự khác biệt giữa Functional Component và Class Component, cách truyền dữ liệu và sự kiện.',
    40, 1, 1, NOW(), NOW()),
(6,  3, 2, 'Lifecycle và React Hooks',
    'Hiểu vòng đời của React component (mount, update, unmount). useState quản lý state cục bộ. useEffect thay thế lifecycle methods, xử lý side effects. useRef, useCallback, useMemo tối ưu performance.',
    50, 2, 0, NOW(), NOW()),
(7,  4, 2, 'State Management với Redux Toolkit',
    'Khi ứng dụng lớn, quản lý state qua Props trở nên phức tạp. Redux Toolkit đơn giản hóa Redux với createSlice, createAsyncThunk. Tích hợp RTK Query để fetch và cache data từ API.',
    60, 1, 0, NOW(), NOW()),
(8,  4, 2, 'React Router và Navigation',
    'React Router v6 quản lý navigation trong SPA. Tìm hiểu BrowserRouter, Route, Link, useNavigate, useParams. Xây dựng protected routes yêu cầu đăng nhập. Lazy loading với Suspense.',
    45, 2, 0, NOW(), NOW()),

-- Course 3: Node.js
(9,  5, 3, 'Node.js Runtime và Event Loop',
    'Node.js cho phép chạy JavaScript ngoài trình duyệt nhờ V8 engine. Event Loop là cơ chế xử lý async non-blocking I/O. Hiểu Call Stack, Callback Queue, Microtask Queue. Modules system: CommonJS và ES Modules.',
    35, 1, 1, NOW(), NOW()),
(10, 5, 3, 'Xây dựng REST API với Express.js',
    'Express là framework tối giản cho Node.js. Thiết lập server, định nghĩa routes với GET/POST/PUT/DELETE. Middleware chain xử lý request. Body parser, CORS, helmet bảo mật. Tổ chức code theo MVC pattern.',
    55, 2, 0, NOW(), NOW()),
(11, 6, 3, 'Kết nối MySQL với Sequelize ORM',
    'Sequelize là ORM mạnh mẽ cho Node.js hỗ trợ MySQL, PostgreSQL, SQLite. Định nghĩa Model, Association (HasMany, BelongsTo, ManyToMany). Migrations quản lý schema thay đổi. Query builder và raw query.',
    50, 1, 0, NOW(), NOW()),
(12, 6, 3, 'JWT Authentication và phân quyền',
    'Xác thực stateless với JSON Web Token. Access token và Refresh token flow. Middleware kiểm tra JWT trên mỗi request. Role-based access control (RBAC) phân quyền admin/user. Bcrypt hash password an toàn.',
    40, 2, 0, NOW(), NOW()),

-- Course 4: React Native
(13, 7, 4, 'Thiết lập môi trường và dự án đầu tiên',
    'Cài đặt Node.js, JDK, Android Studio, Xcode (macOS). Tạo dự án với React Native CLI và Expo. Hiểu cấu trúc thư mục. Chạy app trên Android Emulator và iOS Simulator. Hot Reload và Fast Refresh.',
    30, 1, 1, NOW(), NOW()),
(14, 7, 4, 'Core Components và Styling',
    'React Native không dùng HTML mà dùng native components: View, Text, Image, TextInput, TouchableOpacity, FlatList, ScrollView. StyleSheet API dùng Flexbox để layout. Platform-specific styling cho iOS và Android.',
    45, 2, 0, NOW(), NOW()),
(15, 8, 4, 'React Navigation — Stack, Tab và Drawer',
    'React Navigation là thư viện navigation phổ biến nhất. Stack Navigator cho màn hình dạng stack. Bottom Tab Navigator cho tab bar. Drawer Navigator cho menu slide. Truyền params giữa màn hình. Deep linking.',
    50, 1, 0, NOW(), NOW()),
(16, 8, 4, 'Gọi API và quản lý State',
    'Fetch API và Axios để gọi REST API. Xử lý loading state và error. AsyncStorage lưu dữ liệu local (token, preferences). Context API và Redux Toolkit quản lý global state. Tích hợp notification với Firebase.',
    55, 2, 0, NOW(), NOW()),

-- Course 5: Docker
(17, 9,  5, 'Docker là gì? Container vs Virtual Machine',
    'Container là công nghệ ảo hóa nhẹ ở mức OS, chia sẻ kernel với host. Khác với VM phải ảo hóa toàn bộ OS. Docker Engine, Docker Hub, Docker CLI. Image là template read-only, Container là instance đang chạy từ Image.',
    25, 1, 1, NOW(), NOW()),
(18, 9,  5, 'Dockerfile, Images và Containers',
    'Dockerfile định nghĩa các bước build image: FROM, RUN, COPY, ADD, WORKDIR, ENV, EXPOSE, CMD, ENTRYPOINT. Multi-stage build tối ưu image size. Docker layer caching tăng tốc build. docker build, run, stop, rm, logs, exec.',
    45, 2, 0, NOW(), NOW()),
(19, 10, 5, 'Kubernetes Architecture và Core Concepts',
    'Kubernetes (K8s) là hệ thống orchestration container. Master node: API Server, Scheduler, Controller Manager, etcd. Worker node: kubelet, kube-proxy, Container Runtime. Pod là đơn vị deploy nhỏ nhất. Namespace cô lập tài nguyên.',
    60, 1, 0, NOW(), NOW()),
(20, 10, 5, 'Deployment, Service và Ingress Controller',
    'Deployment quản lý ReplicaSet, đảm bảo số Pod luôn chạy. Rolling update và rollback không downtime. Service expose Pod ra ngoài: ClusterIP, NodePort, LoadBalancer. Ingress là layer 7 load balancer, routing dựa trên domain và path.',
    65, 2, 0, NOW(), NOW());

-- -------------------------------------------------------
-- 6. Quizzes  (1 per course, attached to first lesson)
-- -------------------------------------------------------
INSERT INTO `quizzes` (`id`, `course_id`, `lesson_id`, `title`, `passing_score`, `created_at`, `updated_at`) VALUES
(1, 1, 1,  'Kiểm tra kiến thức JavaScript ES6+',          70, NOW(), NOW()),
(2, 2, 5,  'Kiểm tra kiến thức ReactJS',                  70, NOW(), NOW()),
(3, 3, 9,  'Kiểm tra kiến thức Node.js & Express',        70, NOW(), NOW()),
(4, 4, 13, 'Kiểm tra kiến thức React Native',             70, NOW(), NOW()),
(5, 5, 17, 'Kiểm tra kiến thức Docker & Kubernetes',      70, NOW(), NOW());

-- -------------------------------------------------------
-- 7. Questions  (4 per quiz = 20 total)
-- -------------------------------------------------------
INSERT INTO `questions` (`id`, `quiz_id`, `question_text`, `option_a`, `option_b`, `option_c`, `option_d`, `correct_option`, `created_at`, `updated_at`) VALUES

-- Quiz 1: JavaScript ES6+
(1,  1, 'Trong JavaScript ES6, cú pháp nào dùng để khai báo biến không thể thay đổi giá trị sau khi gán?',
    'var', 'let', 'const', 'define', 'C', NOW(), NOW()),
(2,  1, 'Arrow function trong ES6 có đặc điểm nào sau đây?',
    'Có đối tượng `arguments` riêng như function thường',
    'Không binding `this` từ context bên ngoài',
    'Bắt buộc phải có từ khóa return',
    'Chỉ có thể dùng một lần duy nhất',
    'B', NOW(), NOW()),
(3,  1, 'Promise trong JavaScript được sinh ra để giải quyết vấn đề gì?',
    'Tăng tốc độ xử lý đồng bộ',
    'Quản lý bộ nhớ tốt hơn',
    'Xử lý bất đồng bộ và tránh callback hell',
    'Khai báo biến toàn cục',
    'C', NOW(), NOW()),
(4,  1, 'Cú pháp destructuring `const { name, age } = user` trong ES6 làm gì?',
    'Xóa thuộc tính name và age khỏi object user',
    'Copy toàn bộ object user sang biến mới',
    'Trích xuất giá trị name và age từ object user vào biến riêng',
    'Merge hai object lại với nhau',
    'C', NOW(), NOW()),

-- Quiz 2: ReactJS
(5,  2, 'React Hook nào dùng để khai báo và quản lý state trong functional component?',
    'useEffect', 'useContext', 'useState', 'useRef', 'C', NOW(), NOW()),
(6,  2, 'Trong React, prop drilling là vấn đề gì?',
    'Thêm props không cần thiết vào component',
    'Truyền props qua nhiều cấp component trung gian không dùng đến',
    'Xóa props khỏi component con',
    'Kiểm tra kiểu dữ liệu của props với PropTypes',
    'B', NOW(), NOW()),
(7,  2, 'useEffect với dependency array rỗng [] sẽ được thực thi khi nào?',
    'Mỗi lần component re-render',
    'Chỉ khi state thay đổi',
    'Đúng một lần sau khi component mount lần đầu',
    'Không bao giờ được gọi',
    'C', NOW(), NOW()),
(8,  2, 'Virtual DOM trong React giúp ích gì cho performance?',
    'Lưu trữ dữ liệu của ứng dụng trong bộ nhớ',
    'So sánh và chỉ cập nhật những phần DOM thực sự thay đổi',
    'Xử lý validation dữ liệu form',
    'Tự động quản lý routing giữa các trang',
    'B', NOW(), NOW()),

-- Quiz 3: Node.js & Express
(9,  3, 'Node.js sử dụng mô hình nào để xử lý I/O hiệu quả?',
    'Multi-threading — mỗi request một thread riêng',
    'Event-driven, non-blocking I/O với single thread',
    'Synchronous blocking — xử lý tuần tự từng request',
    'Process forking — tạo process mới cho mỗi request',
    'B', NOW(), NOW()),
(10, 3, 'Middleware trong Express.js có chức năng gì?',
    'Kết nối trực tiếp đến database',
    'Render file HTML và trả về cho client',
    'Xử lý request trước khi đến route handler cuối cùng',
    'Tạo và quản lý HTTP server',
    'C', NOW(), NOW()),
(11, 3, 'Lệnh `npm init -y` tạo ra file nào trong project?',
    'index.js', 'app.js', 'package.json', '.env', 'C', NOW(), NOW()),
(12, 3, 'Trong JWT (JSON Web Token), phần payload chứa gì?',
    'Secret key dùng để ký token',
    'Chữ ký điện tử (signature)',
    'Claims — dữ liệu như userId, role, exp',
    'Thuật toán mã hóa được sử dụng',
    'C', NOW(), NOW()),

-- Quiz 4: React Native
(13, 4, 'StyleSheet.create() trong React Native có lợi ích gì?',
    'Tạo animation phức tạp cho component',
    'Tối ưu performance bằng cách validate và cache styles',
    'Kết nối với native module của iOS/Android',
    'Xử lý navigation giữa các màn hình',
    'B', NOW(), NOW()),
(14, 4, 'FlatList trong React Native tốt hơn ScrollView cho danh sách dài vì?',
    'API đơn giản hơn, dễ sử dụng hơn',
    'Có thêm animation mặc định',
    'Chỉ render các item đang hiển thị trên màn hình (virtualization)',
    'Hỗ trợ cuộn ngang (horizontal) tốt hơn',
    'C', NOW(), NOW()),
(15, 4, 'AsyncStorage trong React Native được dùng để làm gì?',
    'Gọi REST API bất đồng bộ',
    'Lưu trữ dữ liệu key-value local trên thiết bị',
    'Quản lý global state của ứng dụng',
    'Điều hướng giữa các màn hình',
    'B', NOW(), NOW()),
(16, 4, 'Metro Bundler trong React Native có vai trò gì trong quá trình phát triển?',
    'Quản lý dependencies như npm',
    'Bundle JavaScript code thành file duy nhất để chạy trên thiết bị',
    'Compile native code Swift/Kotlin',
    'Quản lý kết nối database',
    'B', NOW(), NOW()),

-- Quiz 5: Docker & Kubernetes
(17, 5, 'Lệnh nào dùng để xem danh sách tất cả Docker containers đang chạy?',
    'docker ps', 'docker ls --containers', 'docker images', 'docker container show',
    'A', NOW(), NOW()),
(18, 5, 'Trong Dockerfile, sự khác biệt chính giữa lệnh COPY và ADD là gì?',
    'Không có sự khác biệt, hai lệnh giống nhau hoàn toàn',
    'ADD hỗ trợ copy từ URL và tự động giải nén file .tar',
    'COPY thực thi nhanh hơn ADD trong mọi trường hợp',
    'ADD chỉ dùng được để copy thư mục, không copy file',
    'B', NOW(), NOW()),
(19, 5, 'Kubernetes Pod là gì?',
    'Tên gọi của toàn bộ Kubernetes cluster',
    'Đơn vị deploy nhỏ nhất, có thể chứa một hoặc nhiều container',
    'Load balancer phân phối traffic giữa các node',
    'Hệ thống lưu trữ persistent volume',
    'B', NOW(), NOW()),
(20, 5, 'Docker Compose được dùng để làm gì?',
    'Build Docker image từ source code',
    'Quản lý Docker network và firewall',
    'Định nghĩa và chạy ứng dụng multi-container bằng file YAML',
    'Push image lên Docker Hub',
    'C', NOW(), NOW());
