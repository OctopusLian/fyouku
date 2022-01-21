<?php
namespace app\models;

use think\facade\Db;

class Barrage
{
    /**
     * 保存弹幕
     * @param int $episodesId
     * @param int $videoId
     * @param int $uid
     * @param string $content
     * @param int $currentTime   时间*100000
     * @return rs
     */
    public static function save($content, $currentTime, $uid, $episodesId, $videoId){
        $data = [
            'content' => $content, 
            'current_time' => $currentTime, 
            'user_id' => $uid,
            'episodes_id' => $episodesId,
            'video_id' => $videoId,
            'status' => 1,
            'add_time' => time()
        ];
        $rs = Db::table('barrage')->insert($data);
        if($rs){
            return true;
        }
        return false;
    }

    /**
     * 按照时间段获取视频剧集下弹幕
     * @param int $episodesId
     * @return rs
     */
    public static function list($episodesId, $startTime, $endTime){
        $rs = Db::table('barrage')
            ->field('content,current_time')
            ->where([
                ['episodes_id', '=', $episodesId],
                ['status', '=', 1],
                ['current_time', '>=', $startTime],
                ['current_time', '<', $endTime],
            ])
            ->order('current_time', 'asc')
            ->select();
            
        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    
}
