<?php
namespace app\controller;

use app\BaseController;
use app\models\Barrage;
use think\facade\Request;
use think\response\Json;

class BarrageController extends BaseController
{
    /**
     * 保存弹幕
     * @param int $episodesId
     * @param int $uid
     * @param int $videoId
     * @param string $content
     * @param int $currentTime
     * @return \think\response\Json
     */
    public function save()
    {
    	try{
	        //获取视频剧集ID
	        $episodesId = Request::post('episodesId', 0);
	        //获取视频ID
	        $videoId = Request::post('videoId', 0);
	        //获取用户ID
	        $uid = Request::post('uid', 0);
	        //获取内容
	        $content = Request::post('content', '');
	        //获取视频播放当前时间
	        $currentTime = Request::post('currentTime', 1);

	        if(!$content){
	        	throw new \Exception("内容不能为空");
	        }
	        if(!$uid){
	        	throw new \Exception("请先登录");
	        }
	        if(!$episodesId){
	        	throw new \Exception("必须指定评论视频集数");
	        }
	        if(!$videoId){
	        	throw new \Exception("必须指定评论视频");
	        }
	        $rs = Barrage::save($content, $currentTime, $uid, $episodesId, $videoId);
	        if($rs){
	        	returnSuccess("保存成功");
	        }
	        throw new \Exception("保存失败");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 按照时间段获取视频剧集下弹幕
     * @param int $episodesId
     * @param int $startTime
     * @param int $endTime
     * @return \think\response\Json
     */
    public function list()
    {
    	try{
	        //获取视频剧集ID
	        $episodesId = Request::get('episodesId', 0);
	        //获取开始时间
	        $startTime = Request::get('startTime', 0);
	        //获取结束时间
	        $endTime = Request::get('endTime', 0);

	        if(!$episodesId){
	        	throw new \Exception("必须指定评论视频集数");
	        }
	        if(!$startTime){
	        	throw new \Exception("必须指定时间段1");
	        }
	        if(!$endTime){
	        	throw new \Exception("必须指定时间段12");
	        }

	        $rs = Barrage::List($episodesId, $startTime, $endTime);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("保存失败");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

}
