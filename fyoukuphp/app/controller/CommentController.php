<?php
namespace app\controller;

use app\BaseController;
use app\models\User;
use app\models\Comment;
use think\facade\Request;
use think\response\Json;

class CommentController extends BaseController
{
    /**
     * 获取评论列表
     * @param int $episodesId
     * @return \think\response\Json
     */
    public function list()
    {
    	try{
	        //获取视频剧集ID
	        $episodesId = Request::get('episodesId', 0);
	        //获取翻页信息
	        $limit = Request::get('limit', 12);
	        $offset = Request::get('offset', 0);
	        if(!$episodesId){
	        	throw new \Exception("episodesId不能为空");
	        }
	        $rs = Comment::getCommentList($episodesId, $limit, $offset);
	        if($rs){
	        	//循环加载用户信息
	        	foreach($rs['data'] as $key => $val){
	        		$rs['data'][$key]['userinfo'] = User::getUserInfo($val['user_id']);
	        		$rs['data'][$key]['addTimeTitle'] = date("Y-m-d H:i:s", $val['add_time']);
	        		$rs['data'][$key]['praiseCount'] = $val['praise_count'];
	        	}
	        	returnSuccess($rs['data'], $rs['count']);
	        }
	        throw new \Exception("没有相关内容");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

    /**
     * 保存评论
     * @param int $episodesId
     * @param int $uid
     * @param int $videoId
     * @param string $content
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
	        $rs = Comment::save($content, $uid, $episodesId, $videoId);
	        if($rs){
	        	returnSuccess("保存成功");
	        }
	        throw new \Exception("保存失败");
    	} catch (\Exception $e) {
		    returnError(5000, $e->getMessage());
		}
    }

}
