$(function() {
    $('.money').each(function (index, element) {
        money = $(element).text()
        $(element).text(money.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ","))
    });

    $('.product-rating').each(function(index, element) {
        $(element).html(createStarGrp($(element).data("score")))
    })
})

function createStarGrp (nums) {
    var darks = Math.floor(nums);
    var half = Math.round(nums%darks);
    var empty = 5 - darks - half;
    var result = '';
    for (var i = 0; i<darks; i++) {
        result += '<i class="fa fa-star"></i> ';
    }
    for (var i = 0; i<half; i++) {
        result += '<i class="fa fa-star-half-empty"></i> ';
    }
    for (var i = 0; i<empty; i++) {
        result += '<i class="fa fa-star-o"></i> ';
    }
    return result;
}
