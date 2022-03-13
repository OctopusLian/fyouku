<?php
namespace app\models;

use think\facade\Db;

class Video
{
    /**
     * 根据频道ID获取正在热播视频
     * @param int $channelId
     * @param int $limit
     * @return rs
     */
    public static function getChannelHotList($channelId, $limit = 9){
        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where([
                ['channel_id', '=', $channelId],
                ['is_hot', '=', 1],
                ['status', '=', 1],
            ])
            ->order('episodes_update_time', 'desc')
            ->limit($limit)
            ->select();

        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 根据频道下类型ID获取推荐视频
     * @param int $channelId
     * @param int $typeId
     * @param int $limit
     * @return rs
     */
    public static function getChannelTypeRecommend($channelId, $typeId, $limit = 9){
        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where([
                ['channel_id', '=', $channelId],
                ['type_id', '=', $typeId],
                ['is_recommend', '=', 1],
                ['status', '=', 1],
            ])
            ->order('episodes_update_time', 'desc')
            ->limit($limit)
            ->select();

        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 根据频道下地区ID获取推荐视频
     * @param int $channelId
     * @param int $regionId
     * @param int $limit
     * @return rs
     */
    public static function getChannelRegionRecommend($channelId, $regionId, $limit = 9){
        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where([
                ['channel_id', '=', $channelId],
                ['region_id', '=', $regionId],
                ['is_recommend', '=', 1],
                ['status', '=', 1],
            ])
            ->order('episodes_update_time', 'desc')
            ->limit($limit)
            ->select();

        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 频道下根据不同条件和排序方式获取视频信息
     * @param int $channelId
     * @param int $regionId
     * @param int $typeId
     * @param string $end
     * @param string $sort
     * @param int $offset
     * @param int $limit
     * @return rs
     */
    public static function getChannelVideoList($channelId, $regionId, $typeId, $end, $sort, $offset, $limit){
        $where = [
            ['channel_id', '=', $channelId],
            ['status', '=', 1]
        ];
        if($regionId){
            array_push($where, ['region_id', '=', $regionId]);
        }
        if($typeId){
            array_push($where, ['type_id', '=', $typeId]);
        }
        if($end == 'n'){
            array_push($where, ['is_end', '=', 0]);
        } elseif($end == 'y'){
            array_push($where, ['is_end', '=', 1]);
        }

        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where($where);

        if($sort == 'episodesUpdateTime'){
            $rs = $rs->order('episodes_update_time', 'desc');
        } elseif($sort == 'comment'){
            $rs = $rs->order('comment', 'desc');
        } elseif($sort == 'addTime'){
            $rs = $rs->order('add_time', 'desc');
        } else {
            $rs = $rs->order('add_time', 'desc');
        }

        $rs = $rs->limit($offset, $limit)->select();
        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 根据用户ID获取用户相关数据
     * @param int $uid
     * @param int $limit
     * @return rs
     */
    public static function getUserVideo($uid){
        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where([
                ['user_id', '=', $uid]
            ])
            ->order('add_time', 'desc')
            ->select();

        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 根据视频ID获取视频信息
     * @param int $videoId
     * @return rs
     */
    public static function getVideoInfo($videoId){
        $rs = Db::table('video')->where([
            ['id', '=', $videoId]
        ])
        ->limit(1)
        ->select();
        $data = [];
        if(!$rs->isEmpty()){
            return $rs[0];
        }
        return false;
    }

    /**
     * 根据视频ID获取剧集列表
     * @param int $videoId
     * @return rs
     */
    public static function getVideoEpisodesList($videoId){
        $rs = Db::table('video_episodes')
            ->field('id,title,add_time,num,play_url as playUrl,comment')
            ->where([
                ['video_id', '=', $videoId]
            ])
            ->order('num', 'asc')
            ->select();

        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }

    /**
     * 搜索视频
     * @param int $uid
     * @param int $limit
     * @return rs
     */
    public static function search($keyword, $offset, $limit){
        $rs = Db::table('video')
            ->field('id,title,sub_title,img,img1,add_time,episodes_count,is_end')
            ->where('title', 'like', '%'.$keyword.'%')
            ->order('add_time', 'desc')
            ->limit($offset, $limit)
            ->select();

        $data = [];
        if(!$rs->isEmpty()){
            $data = $rs;
        }
        return ['count' => count($data), 'data' => $data];
    }
}
