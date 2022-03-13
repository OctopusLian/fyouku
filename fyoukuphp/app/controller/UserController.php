<?php
namespace app\controller;

use app\BaseController;
use app\models\User;
use think\facade\Request;
use think\response\Json;

class UserController extends BaseController
{
    /**
     * 通过手机号和密码登录
     * @param string $mobile
     * @param string $password
     * @return \think\response\Json
     */
    public function loginDo()
    {
    	try{
	        //获取手机号
	        //获取密码
	        $mobile = Request::post('mobile', '');
	        $password = Request::post('password', '');
	        if(!$mobile){
	        	throw new \Exception("手机号不能为空");
	        }
	        if(!$password){
	        	throw new \Exception("密码不能为空");
	        }
	        if(!isMobile($mobile)){
	        	throw new \Exception("手机号格式不正确");
	        }
	        $rs = User::isMobileLogin($mobile, Md5V($password));
	        if($rs){
	        	returnSuccess([
	        		'uid' => $rs['id'], 
	        		'name' => $rs['name']
	        	]);
	        }
	        throw new \Exception("手机号或密码不正确");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 注册用户
     * @param string $mobile
     * @param string $password
     * @return \think\response\Json
     */
    public function saveRegister(){
    	try{
	        //获取手机号
	        //获取密码
	        $mobile = Request::post('mobile', '');
	        $password = Request::post('password', '');
	        if(!$mobile){
	        	throw new \Exception("手机号不能为空");
	        }
	        if(!$password){
	        	throw new \Exception("密码不能为空");
	        }
	        if(!isMobile($mobile)){
	        	throw new \Exception("手机号格式不正确");
	        }
	        $rs = User::isUserMobile($mobile);
	        if($rs){
	        	throw new \Exception("手机号已经注册");
	        } else {
	        	$rs = User::userSave($mobile, Md5V($password));
	        	if($rs){
	        		returnSuccess('注册成功');
	        	}
	        }
	        throw new \Exception("注册失败，请联系客服");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 发送通知消息
     * @param string $uids
     * @param string $content
     * @return \think\response\Json
     */
    public function sendMessageDo(){
    	try{
	        //获取用户ID
	        //获取内容
	        $uids = Request::post('uids', '');
	        $content = Request::post('content', '');
	        if(!$uids){
	        	throw new \Exception("接收人不能为空");
	        }
	        if(!$content){
	        	throw new \Exception("发送内容不能为空");
	        }
	        $uid_array = explode(',', $uids);
	        
	        $message_id = User::sendMessageDo($content);
	        if($message_id){
	        	foreach ($uid_array as $key => $value) {
	        		User::sendMessageUser($value, $message_id);
	        	}
	        	returnSuccess('发送完成');
	        }
	        throw new \Exception("发送失败，请联系客服");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

}
