<?php
// +----------------------------------------------------------------------
// | ThinkPHP [ WE CAN DO IT JUST THINK ]
// +----------------------------------------------------------------------
// | Copyright (c) 2006~2018 http://thinkphp.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: liu21st <liu21st@gmail.com>
// +----------------------------------------------------------------------
use think\facade\Route;

Route::get('think', function () {
    return 'hello,ThinkPHP6!';
});

Route::get('hello/:name', 'index/hello');

//登录接口
Route::post('login/do', 'UserController/loginDo');
//注册接口
Route::post('register/save', 'UserController/saveRegister');

//获取评论列表
Route::get('comment/list', 'CommentController/list');
//保存评论
Route::post('comment/save', 'CommentController/save');

//获取频道下地区信息
Route::get('channel/region', 'BasicsController/channelRegion');
//获取频道下类型信息
Route::get('channel/type', 'BasicsController/channelType');

//频道排行榜
Route::get('channel/top', 'TopController/channelTop');
//类型排行榜
Route::get('type/top', 'TopController/typeTop');

//频道页-顶部广告接口
Route::get('channel/advert', 'VideoController/channelAdvert');
//频道页-正在热播接口
Route::get('channel/hot', 'VideoController/channelHotList');
//频道页-按照地区获取推荐接口
Route::get('channel/recommend/region', 'VideoController/channelRegionRecommendList');
//频道页-按照类型获取推荐接口
Route::get('channel/recommend/type', 'VideoController/channelTypeRecommendList');
//视频列表接口
Route::get('channel/video', 'VideoController/channelVideo');
//搜索接口
Route::post('video/search', 'VideoController/search');
//我的视频接口
Route::get('user/video', 'VideoController/userVideo');
//视频详情
Route::get('video/info', 'VideoController/videoInfo');
//视频剧集列表
Route::get('video/episodes/list', 'VideoController/videoEpisodesList');

//发送消息
Route::post('send/message', 'UserController/sendMessageDo');

//发送弹幕
Route::post('barrage/save', 'BarrageController/save');
//按照时间段获取弹幕列表
Route::get('barrage/list', 'BarrageController/list');

