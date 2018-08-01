var observe;
if (window.attachEvent) {
	observe = function (element, event, handler) {
		element.attachEvent('on'+event, handler);
	};
}
else {
	observe = function (element, event, handler) {
		element.addEventListener(event, handler, false);
	};
}
function init (id) {
	var text = document.getElementById(id);
	function resize () {
		text.style.height = 'auto';
		text.style.height = text.scrollHeight+'px';
	}
	/* 0-timeout to get the already changed text */
	function delayedResize () {
		window.setTimeout(resize, 0);
	}
	observe(text, 'change',  resize);
	observe(text, 'cut',     delayedResize);
	observe(text, 'paste',   delayedResize);
	observe(text, 'drop',    delayedResize);
	observe(text, 'keydown', delayedResize);

	text.focus();
	text.select();
	resize();
}
$(document).ready(function(){
	init('base');
	init('expo');
	init('modulous');
});

function validateNumber(ta)
{
	var num = ta.val().trim();
	if(isNaN(num))
	{
		ta.next().show();
		throw "Not a number";
	}
}

function getModExpoValue(a, b, m)
{
	var res = bigInt.one;
	while(b.isZero() == false)
	{
		if(b.isEven() == false)
			res = res.multiply(a).mod(m);
		a = a.square().mod(m);
		b = b.shiftRight(1);
	}
	return res.toString();
}

function calculateModularExponentiation()
{
	try
	{
		validateNumber($('#base'));
		validateNumber($('#expo'));
		validateNumber($('#modulous'));
		var a = $('#base').val().trim();
		var b = $('#expo').val().trim();
		var m = $('#modulous').val().trim();

		$('#mexpo').val(getModExpoValue(bigInt(a), bigInt(b), bigInt(m)));
		init('mexpo');
	}
	catch (err)
	{}
}
