#!/bin/bash

WORKDIR=`pwd`

CLUSTER=$1
ENV=$2
CMD=$3
JUMPTEST=$4
USECGO=$5
CMDENV=${6/&/"\&"}

clusterArr=(`echo $CLUSTER | sed 's/,/\n/g'`)
envArr=(`echo $ENV | sed 's/,/\n/g'`)

# 仅发布到唯一集群中
if [ ${#clusterArr[@]} -eq 1 ]; then
    if [ ! -d "$WORKDIR/build/$CLUSTER" ]; then
        echo "error: cluster '$CLUSTER' not exists" >&2
        exit 1
    fi

    # 仅发布到唯一环境中
    if [ ${#envArr[@]} -eq 1 ]; then
        if [ ! -d "$WORKDIR/build/$CLUSTER/$ENV" ]; then
            echo "error: environment '$CLUSTER/$ENV' not exists" >&2
            exit 1
        fi

      echo "===== single env and single cluster ====="
      curl --header "PRIVATE-TOKEN:gpKth4Rswoh8vh3hsDU2" https://git.jiaxianghudong.com/api/v4/projects/621/repository/files/scripts%2F.gitlab-ci-template.yml/raw?ref=master | sed s/'{{TAGS}}'/$ENV/g | sed s/'{{CLUSTER}}'/"$CLUSTER"/g | sed s/'{{ENV}}'/"$ENV"/g | sed s/'{{CMD}}'/"$CMD"/g | sed s/'{{JUMPTEST}}'/"$JUMPTEST"/g | sed s/'{{USECGO}}'/"$USECGO"/g | sed s/'{{CMDENV}}'/"$CMDENV"/g > .gitlab-ci-complete.yml
      exit 0
    fi
fi

# 需发布到多个环境中
curl --header "PRIVATE-TOKEN:gpKth4Rswoh8vh3hsDU2" https://git.jiaxianghudong.com/api/v4/projects/621/repository/files/scripts%2F.gitlab-ci-template-for-multi-cluster-env.yml/raw?ref=master | sed s/'{{{TAGS}}}'/$ENV/g | sed s/'{{{CLUSTER}}}'/"$CLUSTER"/g | sed s/'{{{ENV}}}'/"$ENV"/g | sed s/'{{{CMD}}}'/"$CMD"/g | sed s/'{{{JUMPTEST}}}'/"$JUMPTEST"/g | sed s/'{{{USECGO}}}'/"$USECGO"/g | sed s/'{{{CMDENV}}}'/"$CMDENV"/g > .gitlab-ci-complete.yml

echo "Raw: .gitlab-ci-complete.yml"
cat .gitlab-ci-complete.yml

artifacts=""

for ecluster in ${clusterArr[*]}; do
for envm in ${envArr[*]}; do

    if [ ! -d "$WORKDIR/build/$ecluster/$envm" ]; then
        echo "error: environment '$ecluster/$envm' not exists" >&2
        exit 1
    fi

    artifacts="$artifacts\n            - .gitlab-ci-complete-$ecluster-$envm.yml"

cat << EOF >> .gitlab-ci-complete.yml


trigger_pipeline_${ecluster}_${envm}:
    stage: triggers
    trigger:
        strategy: depend
        include:
            - artifact: .gitlab-ci-complete-$ecluster-$envm.yml
              job: generate_config
EOF

done
done

sed -i "s/{{{artifacts}}}/$artifacts/g" .gitlab-ci-complete.yml