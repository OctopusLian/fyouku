<?php
namespace app\models;

use think\facade\Db;

class Basics
{
    /**
     * 获取频道下地区
     * @param int $channelId
     * @return rs
     */
    public static function getChannelRegion($channelId){
        $rs = Db::table('channel_region')
        ->field('id,name')
        ->where([
            ['channel_id', '=', $channelId],
            ['status', '=', 1],
        ])
        ->select();
        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 获取频道下类型
     * @param int $channelId
     * @return rs
     */
    public static function getChannelType($channelId){
        $rs = Db::table('channel_type')
        ->field('id,name')
        ->where([
            ['channel_id', '=', $channelId],
            ['status', '=', 1],
        ])
        ->select();
        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }
}
