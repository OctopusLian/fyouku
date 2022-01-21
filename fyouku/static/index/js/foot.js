$(function(){
    var uid = localStorage.getItem("uid");
	if(uid){
		$(".u-vip_3-BwE").hide();
		$(".a-avatar_17DKW").attr('src', localStorage.getItem("userico"));
	}
});