<?php
namespace app\controller;

use app\BaseController;
use app\models\Top;
use think\facade\Request;
use think\response\Json;

class TopController extends BaseController
{
    /**
     * 根据频道ID获取排行榜
     * @param int $channelId
     * @return \think\response\Json
     */
    public function channelTop()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', '');
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }
	        $rs = Top::getChannelTop($channelId);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 根据类型ID获取排行榜
     * @param int $typeId
     * @return \think\response\Json
     */
    public function typeTop()
    {
    	try{
	        //获取类型ID
	        $typeId = Request::get('typeId', '');
	        if(!$typeId){
	        	throw new \Exception("必须指定类型");
	        }
	        $rs = Top::getTypeTop($typeId);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

}
