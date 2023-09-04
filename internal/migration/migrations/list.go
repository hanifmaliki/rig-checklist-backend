package migrations

import "github.com/go-gormigrate/gormigrate/v2"

var Migrations = []*gormigrate.Migration{
	&mig20230811124400_user,
	&mig20230811124401_position,
	&mig20230811124402_rig,
	&mig20230811124403_location,
	&mig20230811124404_well,
	&mig20230811124405_field,
	&mig20230811124406_activity,
	&mig20230811124407_section,
	&mig20230811124408_sub_section,
	&mig20230811124409_classification,
	&mig20230811124410_question,
	&mig20230811124411_form,
	&mig20230811124412_answer,
	&mig20230811124413_photo,
	&mig20230811124414_sign,
}
