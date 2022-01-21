<?php
namespace app\controller;

use app\BaseController;
use app\models\Video;
use app\models\Advert;
use think\facade\Request;
use think\response\Json;

class VideoController extends BaseController
{
    /**
     * 获取频道顶部广告
     * @param int $channelId
     * @return \think\response\Json
     */
    public function channelAdvert()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', 0);
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }
	        $rs = Advert::getChannelAdvert($channelId);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 获取正在热播列表
     * @param int $channelId
     * @return \think\response\Json
     */
    public function channelHotList()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', 0);
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }

	        $rs = Video::getChannelHotList($channelId, 9);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 获取日漫、国漫推荐
     * @param int $channelId
     * @param int $regionId
     * @return \think\response\Json
     */
    public function channelRegionRecommendList()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', 0);
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }
	        //获取日漫、国漫ID
	        $regionId = Request::get('regionId', 0);
	        if(!$regionId){
	        	throw new \Exception("必须指定频道地区");
	        }

	        $rs = Video::getChannelRegionRecommend($channelId, $regionId, 9);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 获取少女推荐
     * @param int $channelId
     * @param int $typeId
     * @return \think\response\Json
     */
    public function channelTypeRecommendList()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', 0);
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }
	        //获取少女类型ID
	        $typeId = Request::get('typeId', 0);
	        if(!$typeId){
	        	throw new \Exception("必须指定频道类型");
	        }

	        $rs = Video::getChannelTypeRecommend($channelId, $typeId, 9);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 频道页，根据接收到的参数获取视频信息
     * @param int $channelId
     * @param int $typeId
     * @return \think\response\Json
     */
    public function ChannelVideo()
    {
    	try{
	        //获取频道ID
	        $channelId = Request::get('channelId', 0);
	        if(!$channelId){
	        	throw new \Exception("必须指定频道");
	        }
	        //获取地区ID
	        $regionId = Request::get('regionId', 0);
	        //获取类型ID
	        $typeId = Request::get('typeId', 0);
	        //获取状态
	        $end = Request::get('end', '');
	        //获取排序
	        $sort = Request::get('sort', '');

	        $limit = Request::get('limit', 12);
	        $offset = Request::get('offset', 0);

	        $rs = Video::getChannelVideoList($channelId, $regionId, $typeId, $end, $sort, $offset, $limit);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 根据用户ID获取用户相关数据
     * @param int $uid
     * @return \think\response\Json
     */
    public function userVideo()
    {
    	try{
	        //获取用户ID
	        $uid = Request::get('uid', 0);
	        if(!$uid){
	        	throw new \Exception("必须指定用户");
	        }

	        $rs = Video::getUserVideo($uid, 9);
	        if($rs){
	        	$data = [];
	        	foreach ($rs['data'] as $key => $value) {
	        		array_push($data, [
	        			'Id' => $value['id'],
	        			'Img' => $value['img'],
	        			'Title' => $value['title'],
	        			'AddTime' => $value['add_time'],
	        		]);
	        	}
	        	returnSuccess($data, $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 获取视频详情
     * @param int $videoId
     * @return \think\response\Json
     */
    public function videoInfo()
    {
    	try{
	        //获取视频ID
	        $videoId = Request::get('videoId', 0);
	        if(!$videoId){
	        	throw new \Exception("必须指定视频ID");
	        }

	        $rs = Video::getVideoInfo($videoId);
	        if($rs){
	        	returnSuccess($rs);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 获取视频剧集信息
     * @param int $videoId
     * @return \think\response\Json
     */
    public function videoEpisodesList()
    {
    	try{
	        //获取视频ID
	        $videoId = Request::get('videoId', 0);
	        if(!$videoId){
	        	throw new \Exception("必须指定视频ID");
	        }

	        $rs = Video::getVideoEpisodesList($videoId);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 获取视频剧集信息
     * @param string $keyword
     * @return \think\response\Json
     */
    public function search()
    {
    	try{
	        //获取视频ID
	        $keyword = Request::post('keyword', '');
	        if(!$keyword){
	        	throw new \Exception("请输入搜索关键字");
	        }
	        $limit = Request::post('limit', 12);
	        $offset = Request::post('offset', 0);

	        $rs = Video::search($keyword, $offset, $limit);
	        if($rs){
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

}
