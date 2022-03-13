<?php
namespace app\models;

use think\facade\Db;

class User
{
    /**
     * 根据手机号和密码获取数据
     * @param string $mobile
     * @param string $password
     * @return rs
     */
    public static function isMobileLogin($mobile, $password){
    	$rs = Db::table('user')->where([
	        ['mobile', '=', $mobile],
	        ['password', '=', $password],
	    ])
	    ->select();
        if(!$rs->isEmpty()){
            return $rs[0];
        }
	    return false;
    }

    /**
     * 根据手机号获取数据
     * @param string $mobile
     * @return rs
     */
    public static function isUserMobile($mobile){
        $rs = Db::table('user')->where([
            ['mobile', '=', $mobile],
        ])
        ->select();
        if(!$rs->isEmpty()){
            return $rs[0];
        }
        return false;
    }

    /**
     * 保存用户
     * @param string $mobile
     * @param string $password
     * @return rs
     */
    public static function userSave($mobile, $password){
        $data = [
            'name' => '', 
            'mobile' => $mobile,
            'password' => $password,
            'status' => 1,
            'add_time' => time()
        ];
        $rs = Db::table('user')->insert($data);
        if($rs){
            return true;
        }
        return false;
    }

    /**
     * 根据用户ID获取用户信息
     * @param int $id
     * @return rs
     */
    public static function getUserInfo($id){
        $rs = Db::table('user')
        ->field('id,name,add_time,avatar')
        ->where([
            ['id', '=', $id],
        ])
        ->select();
        if(!$rs->isEmpty()){
            return $rs[0];
        }
        return false;
    }

    /**
     * 发送消息
     * @param string $content
     * @return rs
     */
    public static function sendMessageDo($content){
        $data = [
            'content' => $content, 
            'add_time' => time()
        ];
        $message_id = Db::table('message')->insertGetId($data);
        if($message_id){
            return $message_id;
        }
        return false;
    }

    /**
     * 保存消息所属用户
     * @param int $user_id
     * @param int $message_id
     * @return rs
     */
    public static function sendMessageUser($user_id, $message_id){
        $data = [
            'user_id' => $user_id, 
            'message_id' => $message_id,
            'status' => 1,
            'add_time' => time()
        ];
        $rs = Db::table('message_user')->insert($data);
        if($rs){
            return true;
        }
        return false;
    }
}
