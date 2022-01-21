<?php
namespace app\models;

use think\facade\Db;

class Advert
{
    /**
     * 根据频道ID获取顶部广告
     * @param int $channelId
     * @return rs
     */
    public static function getChannelAdvert($channelId){
        $rs = Db::table('advert')
            ->field('id,title,sub_title,img,add_time,url')
            ->where([
                ['channel_id', '=', $channelId],
                ['status', '=', 1],
            ])
            ->order('sort', 'desc')
            ->limit(1)
            ->select();
            
        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }
}
