<?php
namespace app\models;

use think\facade\Db;

class Top
{
    /**
     * 获取频道下排行榜
     * @param int $channelId
     * @return rs
     */
    public static function getChannelTop($channelId){
        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where([
                ['channel_id', '=', $channelId],
                ['status', '=', 1],
            ])
            ->order('comment', 'desc')
            ->limit(10)
            ->select();
        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 获取类型下排行榜
     * @param int $typeId
     * @return rs
     */
    public static function getTypeTop($typeId){
        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where([
                ['type_id', '=', $typeId],
                ['status', '=', 1],
            ])
            ->order('comment', 'desc')
            ->limit(10)
            ->select();
        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }
}
