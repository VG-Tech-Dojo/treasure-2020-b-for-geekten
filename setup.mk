# treasure-app から各チームリポジトリ生成時に変更必要な部分をまとめてやる用
# 普段は使わないはず
# gsed 4.2.2+ が必要 -> brew install gnu-sed
GSED := gsed

GROUP_NAME := b
ACCOUNT_ID := 281032752072

setup:
	find .github -type f -print0 | xargs -0 $(GSED) -i -e "s/treasure-2020-x/treasure-2020-$(GROUP_NAME)/g"
	find . -type f -name "task_def*" -print0 | xargs -0 $(GSED) -i -e "s/treasure-2020-x/treasure-2020-$(GROUP_NAME)/g"
	$(GSED) -i -e "s/group-x/group-$(GROUP_NAME)/g" task_def.frontend.json
	find . -type f -name "task_def*" -print0 | xargs -0 $(GSED) -i -e "s/030993013703/$(ACCOUNT_ID)/g"
	$(GSED) -i -e "s/VG-Tech-Dojo\/treasure-app/VG-Tech-Dojo\/treasure-2020-$(GROUP_NAME)/g" README.md
	find ./backend -type f -print0 | xargs -0 $(GSED) -i -e "s/treasure-app/treasure-2020-$(GROUP_NAME)/g"

