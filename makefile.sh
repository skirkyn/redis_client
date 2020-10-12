#!/bin/bash
env_vars_options="flaxxed_ds_projectId=flaxxed,flaxxed_redis_url=172.23.144.3:6379,flaxxed_menu_pub_sub=flaxxed-menu"

curr_dir=$(pwd)
project_name=flaxxed
pkg_dir=$curr_dir/pkg
dist_dir=$curr_dir/dist
config_file=$curr_dir/configs/config.json
deploy=$1
[[ -d $dist_dir ]] && rm -rf $dist_dir || mkdir -p $dist_dir
modules=($(ls $curr_dir/pkg))
function_modules=($(ls $curr_dir/pkg/function))
#function_modules=(menu)
for mod in "${function_modules[@]}"; do
  lower_func=$(echo $mod | tr "[:upper:]" "[:lower:]")
  . $curr_dir/pkg/function/${mod}/.function

  mkdir -p $dist_dir/$lower_func/vendor/$project_name/pkg

  cp $pkg_dir/function/$mod/* $dist_dir/$lower_func

  cd $dist_dir/$lower_func
  wire
  mv -f wire_gen.go function.go
  go mod init
  go mod tidy
  go mod vendor
  rm go.mod
  rm go.sum
  echo $modules
  for m in "${modules[@]}"; do
     [[ $m != "function" ]] && mkdir -p $dist_dir/$lower_func/vendor/$project_name/pkg/$m && cp -r $pkg_dir/$m/* $dist_dir/$lower_func/vendor/$project_name/pkg/$m
  done

  [[ $deploy ]] && gcloud functions deploy $lower_func --entry-point $function_name --runtime go113 --trigger-${trigger} --allow-unauthenticated --timeout=60 --vpc-connector projects/flaxxed/locations/us-central1/connectors/redis-connector --memory=${memory}MB --set-env-vars $env_vars_options &
  cd $curr_dir
  #    mkdir -p $dist_dir/$lower_func/vendor/$project_name/internal
done
