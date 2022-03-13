<?php
	// $a = 'Hello World';
	// echo $a;
	// var_dump($a);

	// $data = [
	// 	'a',
	// 	'b',
	// 	'c'
	// ];

	// $data1 = [
	// 	'a' => '张三',
	// 	'b' => '李四',
	// 	'c' => '王五'
	// ];

	// var_dump($data1);
	// var_dump($data[1]);
	// var_dump($data1['b']);

	// function showHelloWorld(){
	// 	$a = 'Hello';
	// 	$b = 'World';
	// 	$c = $a . ' ' . $b;
	// 	return $c;
	// }

	// var_dump(showHelloWorld());

	// function showUserInfo(){
	// 	$username = $_GET['username'];
	// 	$password = $_GET['password'];
	// 	return '用户名是：' . $username . '，密码是：' . $password;
	// }

	// var_dump(showUserInfo());

	// class User{
	// 	public function __construct(){
	// 		//这里是构造函数
	// 	}

	// 	public function getUserInfo(){

	// 	}
	// }

	//连接数据库
$conn = new mysqli('127.0.0.1', 'root', '123456', 'fyouku');
if($conn->connect_error){
	die("连接失败：" . $conn->connect_error);
}
$conn->query("SET NAMES utf8");
//var_dump($conn);

$sql = "SELECT * FROM user WHERE id>10 ORDER BY id DESC LIMIT 2";
$result = $conn->query($sql);
if($result->num_rows > 0){
	//输出数据
	while($row = $result->fetch_assoc()){
		echo "用户名：" . $row['name'] . "手机号：" . $row['mobile'];
		echo "<br>";
	}
}

$sql = "INSERT INTO user (name, mobile, status, add_time, avatar) VALUES ('test3', '18600008888', 0, '".time()."', '我是头像')";
if($conn->query($sql) == true){
	echo "保存用户成功";
} else {
	echo "error:" . $conn->error;
}


$conn->close();














?>
