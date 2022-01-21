<?php
namespace app\models;

use think\facade\Db;

class Comment
{
    /**
     * 根据剧集数获取评论列表
     * @param int $episodesId
     * @param int $limit
     * @param int $offset
     * @return rs
     */
    public static function getCommentList($episodesId, $limit, $offset){
    	$rs = Db::table('comment')->where([
	        ['episodes_id', '=', $episodesId],
	        ['status', '=', 1],
	    ]);
        $count = $rs->count();
	    $rs = $rs->limit($offset,$limit)->order('add_time', 'desc')->select()->all();
        $data = [];
        if(count($rs)){
            $data = $rs;
        }
	    return ['count' => $count, 'data' => $data];
    }

    /**
     * 保存评论
     * @param int $episodesId
     * @param int $videoId
     * @param int $uid
     * @param string $content
     * @return rs
     */
    public static function save($content, $uid, $episodesId, $videoId){
        $data = [
            'content' => $content, 
            'user_id' => $uid,
            'episodes_id' => $episodesId,
            'video_id' => $videoId,
            'stamp' => 0,
            'status' => 1,
            'add_time' => time()
        ];
        $rs = Db::table('comment')->insert($data);
        if($rs){
            //更新视频总评论数
            Db::table('video')->where('id', $videoId)->inc('comment')->update();
            //更新视频剧集评论数
            Db::table('video_episodes')->where('id', $episodesId)->inc('comment')->update();
            return true;
        }
        return false;
    }
}
