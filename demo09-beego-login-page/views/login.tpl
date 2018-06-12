<!DOCTYPE html>
<html>
<head>
    <title>login</title>
    <meta charset="utf-8">
    <meta http-equiv="x-ua-compatible" content="ie=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1"></head>
    <link type="text/css" rel="stylesheet" href="/admin/assets/stylesheets/qor_admin_default.css">
<body>

<header class="mdl-layout__header is-casting-shadow has-search"><div aria-expanded="false" role="button" tabindex="0" class="mdl-layout__drawer-button"><i class="material-icons"></i></div>
        <div class="mdl-layout__header-row">


            <span class="mdl-layout-title">Asset Managers</span>



    <div class="mdl-layout-spacer"></div>
    <button class="mdl-button mdl-navigation mdl-layout--small-screen-only qor-mobile--show-actions">Actions <i class="material-icons">arrow_drop_down</i></button>




    <form class="qor-search-container ignore-dirtyform" method="GET">



            <input name="order_by" value="file" type="hidden">




      <div class="mdl-textfield mdl-js-textfield mdl-textfield--expandable qor-search has-placeholder is-upgraded" data-upgraded=",MaterialTextfield">
        <label class="mdl-button mdl-js-button mdl-button--icon qor-search__label" for="inputSearch" data-upgraded=",MaterialButton">
          <i class="material-icons">search</i>
        </label>
        <div class="mdl-textfield__expandable-holder">
          <input class="mdl-textfield__input qor-search__input" type="text" id="inputSearch" name="keyword" value="" placeholder="Search">
          <label class="mdl-textfield__label"></label>
        </div>
        <button class="mdl-button mdl-js-button mdl-button--icon mdl-button--colored qor-search__clear" type="button" data-upgraded=",MaterialButton">
          <i class="material-icons md-18">clear</i>
        </button>
      </div>
    </form>

        </div>
      </header>
  <form action="/auth/login" method="POST">
    UserName:    <input name="UserName">
    Password: <input name="Password" type="password">
    <input type="submit" class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored">
  </form>

  <div class="qor-form-container mdl-layout__content" style="width:300px;text-align:center">
      <form class="qor-form" action="/admin/user_infos" method="POST" enctype="multipart/form-data">
        <div class="qor-form-section clearfix" data-section-title="">
        <div>
        <div class="qor-form-section-rows qor-section-columns-1 clearfix">
          <input id="" class="qor-hidden__primary_key" name="QorResource.ID" value="0" type="hidden">
        </div>
    </div>
  </div>
  <div class="qor-form-section clearfix" data-section-title="">
    <div>
        <div class="qor-form-section-rows qor-section-columns-1 clearfix">
          <div class="qor-field">
        <div class="mdl-textfield mdl-textfield--full-width mdl-js-textfield is-upgraded" data-upgraded=",MaterialTextfield">
        <label class="qor-field__label mdl-textfield__label" for="">
        User Name
      </label>

      <div class="qor-field__show"></div>

      <div class="qor-field__edit">
        <input class="mdl-textfield__input" type="text" id="" name="QorResource.UserName" value="">
      </div>
    </div>
  </div>

        </div>

    </div>
  </div>



          <div class="qor-form__actions">
            <button class="mdl-button mdl-button--colored mdl-button--raised mdl-js-button mdl-js-ripple-effect qor-button--save" type="submit" data-upgraded=",MaterialButton,MaterialRipple">Add<span class="mdl-button__ripple-container"><span class="mdl-ripple"></span></span></button>

            <a class="mdl-button mdl-button--primary mdl-js-button mdl-js-ripple-effect qor-button--cancel" data-dismiss="slideout" data-upgraded=",MaterialButton,MaterialRipple">Cancel<span class="mdl-button__ripple-container"><span class="mdl-ripple"></span></span></a>
          </div>

      </form>
    </div>
  <div style="color: red;">
        {{if .Err}}
            用户名密码错误
        {{else}}
        {{end}}
  </div>


</body>
</html>
