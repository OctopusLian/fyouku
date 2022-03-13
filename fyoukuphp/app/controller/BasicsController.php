<?php
namespace app\controller;

use app\BaseController;
use app\models\Basics;
use think\facade\Request;
use think\response\Json;

class BasicsController extends BaseController
{
    /**
     * 获取频道下地区
     * @param int $channelId
     * @return \think\response\Json
     */
    public function channelRegion()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', 0);
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }
	        $rs = Basics::getChannelRegion($channelId);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 获取频道类型
     * @param int $channelId
     * @return \think\response\Json
     */
    public function channelType()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', 0);
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }
	        $rs = Basics::getChannelType($channelId);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

}
