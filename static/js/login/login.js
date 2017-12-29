$(function() {
	var ok1 = false;
	var ok2 = false;

	// 验证用户名
	$('input[name="username"]').focus(function() {

	}).blur(function() {
		if($(this).val().length >= 3 && $(this).val().length <= 12 && $(this).val() != '') {
			$(this).next().text('输入成功').removeClass('state1').addClass('state4');
			ok1 = true;
		} else {
			$(this).next().show().text('用户名应该为3-20位之间');
		}

	});

	//验证密码
	$('input[name="password"]').focus(function() {

	}).blur(function() {
		if($(this).val().length >= 6 && $(this).val().length <= 20 && $(this).val() != '') {
			$(this).next().text('输入成功');
			ok2 = true;
		} else {
			$(this).next().text('密码应该为6-20位之间');
		}

	});
	//提交按钮,所有验证通过方可提交

	$('input[type="button"]').click(function() {
		if(ok1 && ok2) {
			$('form').submit();
		} else {
			return false;
		}
	});

});