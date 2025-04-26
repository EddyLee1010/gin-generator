package templates

// ConfigFileTemplateStr  生成config.yaml项目运行配置文件的模板
const ConfigFileTemplateStr = `##     ######   #### ##    ##          ######   ######## ##    ## ######## ########     ###    ########  #######  ######## 
##    ##    ##   ##  ###   ##         ##    ##  ##       ###   ## ##       ##     ##   ## ##      ##    ##     ## ##     ##
##    ##         ##  ####  ##         ##        ##       ####  ## ##       ##     ##  ##   ##     ##    ##     ## ##     ##
##    ##   ####  ##  ## ## ## ####### ##   #### ######   ## ## ## ######   ########  ##     ##    ##    ##     ## ######## 
##    ##    ##   ##  ##  ####         ##    ##  ##       ##  #### ##       ##   ##   #########    ##    ##     ## ##   ##  
##    ##    ##   ##  ##   ###         ##    ##  ##       ##   ### ##       ##    ##  ##     ##    ##    ##     ## ##    ## 
##     ######   #### ##    ##         ######   ######## ##    ## ######## ##     ## ##     ##    ##     #######  ##     ##
##
##	Author: Eddy 527084800
app_name: {{.ProjectName}}
port: {{.Port}}
database:
  host: {{.Database.DBHost}}
  port: {{.Database.DBPort}}
  user: {{.Database.DBUser}}
  password: {{.Database.DBPassword}}
  dbname: {{.Database.DBName}}
`
