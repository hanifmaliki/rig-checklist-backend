package seeds

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"

	"gorm.io/gorm"
)

func SeedProduct(db *gorm.DB) error {
	productSeeds := []model.Product{
		{
			Slug:                    "minerva-apps",
			Name:                    "Minerva Apps",
			Desc:                    "Minerva App provides access to real-time performance at the equipment level and personnel status (Production KPI and operator KPI) to accelerate decision making.",
			BannerImageStyle:        `{"maxWidth":"577px","bottom":"-90px","left":"80px","width":"unset"}`,
			HighlightedFeatureDesc:  "Optimizing mining for short-term and long-term activities, by evaluating plan vs actual performance. Sequencing the stages of mining operation for any immediate or long-term decision. Control the movement and displacement unit based on plan vs actual data.",
			IsActive:                true,
			BannerImage:             "product-1-banner-image.webp",
			BannerBackground:        "product-1-banner-background.webp",
			HighlightedFeatureImage: "product-1-highlight-image.webp",
			HighlightedFeatureList: []*model.ProductHighlightedFeature{
				{
					Name:    "Advanced analytics into action",
					Desc:    "Minerva App provides access to real-time performance at the equipment level and personnel status (Production KPI and operator KPI) to accelerate decision making.",
					Image:   "product-1-highlight-icon-1.webp",
					OrderNo: 1,
				},
				{
					Name:    "Working condition compliance",
					Desc:    "Enable visibility on working condition compliance at Mine area to provide optimum productivity and highest utilization",
					Image:   "product-1-highlight-icon-2.webp",
					OrderNo: 2,
				},
				{
					Name:    "Operator Performance",
					Desc:    "Provide realtime operator performance to maintain the performance based on selected key performance parameters",
					Image:   "product-1-highlight-icon-3.webp",
					OrderNo: 3,
				},
				{
					Name:    "E-Learning, E-Library, Idea improvement and Safety",
					Desc:    "Gamification based learning mobile application to underpin change management efforts and maintain learning culture",
					Image:   "product-1-highlight-icon-4.webp",
					OrderNo: 4,
				},
				{
					Name:    "Shift Plan",
					Desc:    "Prepare tactical strategy before running the operation through resource allocation covering manpower allocation and equipment assignment",
					Image:   "product-1-highlight-icon-5.webp",
					OrderNo: 5,
				},
			},
		},
		{
			Slug:                    "fleet-management-system",
			Name:                    "Fleet Management System",
			Desc:                    "Monitor the performance for every site with certain parameter related in real time. Connecting the environment system by system from IOT to our Minerva platform. Improve the performance based on real time data, combine it with mining management.",
			BannerImageStyle:        `{"left":"-120px","width":"unset","height":"410px"}`,
			HighlightedFeatureDesc:  "Monitor the performance for every site with certain parameter related in real time. Connecting the environment system by system from IOT to our Minerva platform. Improve the performance based on real time data, combine it with mining management.",
			IsActive:                true,
			BannerImage:             "product-2-banner-image.webp",
			BannerBackground:        "product-2-banner-background.webp",
			HighlightedFeatureImage: "product-2-highlight-image.webp",
			HighlightedFeatureList: []*model.ProductHighlightedFeature{
				{
					Name:    "Overall Shift Performance",
					Desc:    "Performance tracking for every shift with some of the key parameters such as Production, Digger Productivity, and many others",
					Image:   "product-2-highlight-icon-1.webp",
					OrderNo: 1,
				},
				{
					Name:    "Optimizer Manager",
					Desc:    "Optimize our fleet movement by comparing every last hour performance vs current performance",
					Image:   "product-2-highlight-icon-2.webp",
					OrderNo: 2,
				},
				{
					Name:    "Excavator Fleet Performance",
					Desc:    "Detailed performance for every working excavator performance on every site",
					Image:   "product-2-highlight-icon-3.webp",
					OrderNo: 3,
				},
				{
					Name:    "Operator Compliance",
					Desc:    "Guidance for operator and ensure our operational excellence working seamlessly and comply with our performance parameter",
					Image:   "product-2-highlight-icon-4.webp",
					OrderNo: 4,
				},
			},
		},
		{
			Slug:                    "mine-planning-optimization",
			Name:                    "Mine Planning Optimization",
			Desc:                    "Optimizing mining for short-term and long-term activities, by evaluating plan vs actual performance. Sequencing the stages of mining operation for any immediate or long-term decision. Control the movement and displacement unit based on plan vs actual data.",
			BannerImageStyle:        `{"width":"unset","height":"496px","bottom":"0px"}`,
			HighlightedFeatureDesc:  "Optimizing mining for short-term and long-term activities, by evaluating plan vs actual performance. Sequencing the stages of mining operation for any immediate or long-term decision. Control the movement and displacement unit based on plan vs actual data.",
			IsActive:                true,
			BannerImage:             "product-3-banner-image.webp",
			BannerBackground:        "product-3-banner-background.webp",
			HighlightedFeatureImage: "product-3-highlight-image.webp",
			HighlightedFeatureList: []*model.ProductHighlightedFeature{
				{
					Name:    "MOPAD",
					Desc:    "Mining map with grade/ quality inventory visibility of multiple ore grade position",
					Image:   "product-3-highlight-icon-1.webp",
					OrderNo: 1,
				},
				{
					Name:    "MOCOM",
					Desc:    "Provide information of ROM status capacity and grade availability",
					Image:   "product-3-highlight-icon-2.webp",
					OrderNo: 2,
				},
				{
					Name:    "TOS ASsignment",
					Desc:    "Provide recommendation of fleet position based on grade database and requirements with Advanced analytics concept",
					Image:   "product-3-highlight-icon-3.webp",
					OrderNo: 3,
				},
				{
					Name:    "STMB Management",
					Desc:    "Guidance ore blending by each ROM availability to meet required market",
					Image:   "product-3-highlight-icon-4.webp",
					OrderNo: 4,
				},
			},
		},
		{
			Slug:                    "excellence-crew-management",
			Name:                    "Excellence Crew Management",
			Desc:                    "Capturing performance of every operators in real time integrated with our FMS. Calculating the performance of every operator daily. Informative insight given to see daily performance ranking of every operators involved.",
			BannerImageStyle:        `{"bottom":"-87px","width":"unset","height":"496px"}`,
			HighlightedFeatureDesc:  "Capturing performance of every operators in real time integrated with our FMS. Calculating the performance of every operator daily. Informative insight given to see daily performance ranking of every operators involved.",
			IsActive:                true,
			BannerImage:             "product-4-banner-image.webp",
			BannerBackground:        "product-4-banner-background.webp",
			HighlightedFeatureImage: "product-4-highlight-image.webp",
			HighlightedFeatureList: []*model.ProductHighlightedFeature{
				{
					Name:    "Operator Scoreboard",
					Desc:    "Give information about operator's performance on daily basis",
					Image:   "product-4-highlight-icon-1.webp",
					OrderNo: 1,
				},
				{
					Name:    "Digger Ranking",
					Desc:    "Give a ranking list of Digger operators along with their performance in every site",
					Image:   "product-4-highlight-icon-2.webp",
					OrderNo: 2,
				},
				{
					Name:    "Truck Ranking",
					Desc:    "Give a ranking list of Truck operators along with their performance in every site",
					Image:   "product-4-highlight-icon-3.webp",
					OrderNo: 3,
				},
			},
		},
		{
			Slug:                    "effective-validation-attendance",
			Name:                    "Effective Validation Attendance",
			Desc:                    "Managing crew productivity by digitally allocate to nurture self-driven performance. Real-time Assignment operator by combining simple technology using RFID. Supervise required operator and allocation even before shift change, so that supervisor can focus on other critical work on early of shift.",
			BannerImageStyle:        `{"width":"unset","height":"585px","bottom":"-104px"}`,
			HighlightedFeatureDesc:  "Managing crew productivity by digitally allocate to nurture self-driven performance. Real-time Assignment operator by combining simple technology using RFID. Supervise required operator and allocation even before shift change, so that supervisor can focus on other critical work on early of shift.",
			IsActive:                true,
			BannerImage:             "product-5-banner-image.webp",
			BannerBackground:        "product-5-banner-background.webp",
			HighlightedFeatureImage: "product-5-highlight-image.webp",
			HighlightedFeatureList: []*model.ProductHighlightedFeature{
				{
					Name:    "Operator Management",
					Desc:    "Allocate operator based on their assignment, equipment, along with equipment status, thus it will improve operator utilization.",
					Image:   "product-5-highlight-icon-1.webp",
					OrderNo: 1,
				},
				{
					Name:    "Automatic Attendance Detection",
					Desc:    "Simplify the validation process of the attendance",
					Image:   "product-5-highlight-icon-2.webp",
					OrderNo: 2,
				},
				{
					Name:    "Visibility on Spare Operator",
					Desc:    "Provide recommendation of any available spare operator",
					Image:   "product-5-highlight-icon-3.webp",
					OrderNo: 3,
				},
			},
		},
		{
			Slug:                    "mine-dash",
			Name:                    "Mine-Dash",
			Desc:                    "Capturing performance in each area from Mining to Milling and defining bottlenecks from the process: Breakdown mine output into ore to crusher movement, total mine operation, and to track the volume. Furthermore, to monitor compliance of execution through grade control technology.",
			BannerImageStyle:        `{"bottom":"0px","width":"unset","height":"460px"}`,
			HighlightedFeatureDesc:  "Capturing performance in each area from Mining to Milling and defining bottlenecks from the process: Breakdown mine output into ore to crusher movement, total mine operation, and to track the volume. Furthermore, to monitor compliance of execution through grade control technology.",
			IsActive:                true,
			BannerImage:             "product-6-banner-image.webp",
			BannerBackground:        "product-6-banner-background.webp",
			HighlightedFeatureImage: "product-6-highlight-image.webp",
			HighlightedFeatureList: []*model.ProductHighlightedFeature{
				{
					Name:    "Production Report",
					Desc:    "Provide single truth of data to manage operation performance with current achievement compare target",
					Image:   "product-6-highlight-icon-1.webp",
					OrderNo: 1,
				},
				{
					Name:    "Fuel performance Report",
					Desc:    "Tracking fuel performance and provide monthly and year to date trend",
					Image:   "product-6-highlight-icon-2.webp",
					OrderNo: 2,
				},
				{
					Name:    "Financial performance report",
					Desc:    "Profit and Loss (PnL) allow root cause analysis revenue and cost. ABC costing provide deep dive categorized analysis on value driver tree analysis of EBT.",
					Image:   "product-6-highlight-icon-3.webp",
					OrderNo: 3,
				},
				{
					Name:    "Mineral Mining Key Performance",
					Desc:    "Capturing performance in each area from Mining to Milling and defining bottlenecks from the process. Furthermore, to monitor compliance of execution through grade control technology.",
					Image:   "product-6-highlight-icon-4.webp",
					OrderNo: 4,
				},
			},
		},
		{
			Slug:                    "digital-twin",
			Name:                    "Digital Twin",
			Desc:                    "Virtually represents mine sites to provide improved operation management experience: Combination of orthophoto and orthoimage in a 3D visualization. Incorporating data into our Minerva platform and monitoring the load and haul performance in real time.",
			BannerImageStyle:        `{"height":"496px","bottom":"0px","width":"unset"}`,
			HighlightedFeatureDesc:  "Virtually represents mine sites to provide improved operation management experience: Combination of orthophoto and orthoimage in a 3D visualization. Incorporating data into our Minerva platform and monitoring the load and haul performance in real time.",
			IsActive:                true,
			BannerImage:             "product-7-banner-image.webp",
			BannerBackground:        "product-7-banner-background.webp",
			HighlightedFeatureImage: "product-7-highlight-image.webp",
			HighlightedFeatureList: []*model.ProductHighlightedFeature{
				{
					Name:    "Digital twin Map – 3D visualization",
					Desc:    "Prompt visualization update on mine surface condition and equipment movement.",
					Image:   "product-7-highlight-icon-1.webp",
					OrderNo: 1,
				},
				{
					Name:    "Automatic Mine plan compliance",
					Desc:    "Monitor where you deviated and need to dig to comply to the plan.",
					Image:   "product-7-highlight-icon-2.webp",
					OrderNo: 2,
				},
				{
					Name:    "Hauling fleet heatmap",
					Desc:    "Understand the hauling fleet actual performance against planned performance.",
					Image:   "product-7-highlight-icon-3.webp",
					OrderNo: 3,
				},
				{
					Name:    "Loading equipment status and performance",
					Desc:    "Process performance utilization and availability charts and KPIs.",
					Image:   "product-7-highlight-icon-4.webp",
					OrderNo: 4,
				},
			},
		},
	}

	productMenuSeeds := map[string]model.ProductMenu{
		"Minerva Apps": {
			Type:     "",
			Desc:     "Advanced Analytics Into Action.",
			IsActive: true,
			OrderNo:  1,
		},
		"Fleet Management System": {
			Type:     "Mining dashboard",
			Desc:     "Advanced analytics into action.",
			IsActive: true,
			OrderNo:  2,
		},
		"Mine Planning Optimization": {
			Type:     "Mining dashboard",
			Desc:     "Optimize mining activities",
			IsActive: true,
			OrderNo:  3,
		},
		"Excellence Crew Management": {
			Type:     "Mining dashboard",
			Desc:     "Capture operator performance real time",
			IsActive: true,
			OrderNo:  4,
		},
		"Effective Validation Attendance": {
			Type:     "Mining dashboard",
			Desc:     "Manage crew productivity performance",
			IsActive: true,
			OrderNo:  5,
		},
		"Mine-Dash": {
			Type:     "",
			Desc:     "Capture performance and define bottleneck",
			IsActive: true,
			OrderNo:  6,
		},
		"Digital Twin": {
			Type:     "",
			Desc:     "Real-time Performance plan compliance",
			IsActive: true,
			OrderNo:  7,
		},
	}

	for _, productSeed := range productSeeds {
		productResult, err := service.Instance().CreateProduct(helper.UserDummy, &productSeed)
		if err != nil {
			return err
		}

		productMenuSeed := productMenuSeeds[productResult.Name]
		productMenuSeed.ProductID = productResult.ID

		_, err = service.Instance().CreateProductMenu(helper.UserDummy, &productMenuSeed)
		if err != nil {
			return err
		}
	}

	return nil
}
