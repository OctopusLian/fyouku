<?php
// 应用公共文件

/**
 * json格式返回 - 正确值
 * @param int $code 状态码
 * @param string array $items 返回数据
 * @return json
 */
function returnSuccess($data, $count = 0){
	$return = [
		'code' => 0,
		'items' => $data,
		'count' => $count
	];
	echo json_encode($return);
	exit;
}

/**
 * json格式返回 - 错误值
 * @param int $code 状态码
 * @param string $msg 错误信息
 * @return json
 */
function returnError($code, $msg){
	$return = [
		'code' => $code,
		'msg' => $msg,
	];
	echo json_encode($return);
	exit;
}

/**
 * 判断手机号格式
 * @param string $mobile 手机号
 * @return json
 */
function isMobile($mobile){
    $result = preg_match('b^', $mobile);
    if ($result) {
        return true;
    } else {
        return false;
    }
}

/**
 * 用户密码MD5方法
 * @param string $password 密码
 * @return json
 */
function Md5V($password){
    return md5($password . 'RaW#XhH2aVgo!Iy1');
}
