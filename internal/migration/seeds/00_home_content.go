package seeds

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"gorm.io/gorm"
)

func SeedHomeContent(db *gorm.DB) error {
	seeds := []model.HomeContent{{
		Section: "banner",
		Key:     "title",
		Value:   "Transform your digital experience and unlock possibilities with our cutting-edge Digital Mining.",
		IsJson:  false,
	}, {
		Section: "banner",
		Key:     "desc",
		Value:   "Digitalization promotes efficient operating that brings cost-effectiveness and sustainability, which in turn helps organizations innovate and boost productivity.",
		IsJson:  false,
	}, {
		Section: "banner",
		Key:     "background",
		Value:   "home-banner-background.webm",
		IsJson:  false,
	}, {
		Section: "banner",
		Key:     "image",
		Value:   "home-banner-image.webp",
		IsJson:  false,
	}, {
		Section: "banner",
		Key:     "is_active",
		Value:   "true",
		IsJson:  false,
	}, {
		Section: "company_list",
		Key:     "title",
		Value:   "TRUSTED BY COMPANIES ALL OVER THE WORLD",
		IsJson:  false,
	}, {
		Section: "company_list",
		Key:     "list",
		Value:   `[{"name":"NHM","image":"home-company-list-1.webp","order_no":1},{"name":"Petrosea","image":"home-company-list-2.webp","order_no":2}]`,
		IsJson:  true,
	}, {
		Section: "company_list",
		Key:     "is_active",
		Value:   "true",
		IsJson:  false,
	}, {
		Section: "company_info",
		Key:     "title",
		Value:   "What is minerva?",
		IsJson:  false,
	}, {
		Section: "company_info",
		Key:     "desc",
		Value:   `Minerva Digital Platform is a proprietary digital platform that was designed and developed for the needs of establishing a best-in-class mining operation. By leveraging the combination of real-time data, digital twin, and advanced analytics, our goal is to assist mines to optimize operation and improve performance.`,
		IsJson:  false,
	}, {
		Section: "company_info",
		Key:     "background",
		Value:   "home-company-info-background.webp",
		IsJson:  false,
	}, {
		Section: "company_info",
		Key:     "logo",
		Value:   "home-company-info-logo.webp",
		IsJson:  false,
	}, {
		Section: "company_info",
		Key:     "is_active",
		Value:   "true",
		IsJson:  false,
	}, {
		Section: "minerva_ecosystem",
		Key:     "desc",
		Value:   "All digital level secured by security operation center (SOC) – zero trust program and robust cyber security platform",
		IsJson:  false,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-1-desc",
		Value:   "Data collection from various equipment, data transmission though a reliable connection, and data storage in a secure Data warehouse.",
		IsJson:  false,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-2-desc",
		Value:   "Emerging insights from Data analytics on factors impacting operational performance and utilizing cloud technology for modularity",
		IsJson:  false,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-3-desc",
		Value:   "Deploying advanced analytics algorithms to draw insights and provide recommendations, with enhanced output being the emphasis of the interventions.",
		IsJson:  false,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-4-desc",
		Value:   "Centralized planning control, operation, and performance management across operation functions and mine sites. The next-level productivity optimization is delivered using a digital ecosystem of solutions.",
		IsJson:  false,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-1-features",
		Value:   `[{"name":"Data Warehouse","image":"home-ecosystem-level-1-feature-1.svg","order_no":1},{"name":"Network and connectivity","image":"home-ecosystem-level-1-feature-2.svg","order_no":2},{"name":"Internet of things (IoT)","image":"home-ecosystem-level-1-feature-3.svg","order_no":3},{"name":"Fleet management system","image":"home-ecosystem-level-1-feature-4.svg","order_no":4}]`,
		IsJson:  true,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-2-features",
		Value:   `[{"name":"Data analytics and tools","image":"home-ecosystem-level-2-feature-1.svg","order_no":1},{"name":"Integrated platform","image":"home-ecosystem-level-2-feature-2.svg","order_no":2},{"name":"Cloud server migration","image":"home-ecosystem-level-2-feature-3.svg","order_no":3}]`,
		IsJson:  true,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-3-features",
		Value:   `[{"name":"Digital twin","image":"home-ecosystem-level-3-feature-1.svg","order_no":1},{"name":"Advanced analytics and ML","image":"home-ecosystem-level-3-feature-2.svg","order_no":2}]`,
		IsJson:  true,
	}, {
		Section: "minerva_ecosystem",
		Key:     "level-4-features",
		Value:   `[{"name":"Remote operation center","image":"home-ecosystem-level-4-feature-1.svg","order_no":1}]`,
		IsJson:  true,
	}, {
		Section: "minerva_ecosystem",
		Key:     "is_active",
		Value:   "true",
		IsJson:  false,
	}, {
		Section: "products",
		Key:     "list",
		Value:   `[{"product_id":1,"title":"Minerva Apps","desc":"The Minerva App gives frontline staff access to real-time performance at the equipment level as well as personnel status. Data is easily accessible and visible to those who require it, allowing real-time production to be compared to targets.<br/><br/>Minerva Apps promotes better decision making, enhanced usage, higher productivity, and mine plan compliance.","image":"home-product-1.webp","order_no":1},{"product_id":2,"title":"Mining Dashboard","desc":"A mining dashboard allows dispatchers and control centers to make decisions in real-time while tracking the performance of the entire fleet of equipment.<div style='margin-block:8px 0'><div style='margin-bottom:6px'><span style='font-size:20px;color:#3b8b72;margin-right:8px'>●</span>Fleet Management System</div><div style='margin-bottom:6px'><span style='font-size:20px;color:#3b8b72;margin-right:8px'>●</span>Excellence Crew Management</div><div style='margin-bottom:6px'><span style='font-size:20px;color:#3b8b72;margin-right:8px'>●</span>Mine Planning Optimization</div><div><span style='font-size:20px;color:#3b8b72;margin-right:8px'>●</span>Effective Validation Attendance</div></div>","image":"home-product-2.webp","order_no":2},{"product_id":6,"title":"Mine-Dash","desc":"Single source of truth across the operation with the digital control tower. To preserve operational costs and predictable throughput, it includes production, financial, and fuel analysis.","image":"home-product-3.webp","order_no":3},{"product_id":7,"title":"Digital Twin","desc":"Digital twin is a virtual representation of a mining operation, including the physical assets, equipment, processes, and systems used in the mine extraction process.<br/><br/>The use of Digital twin provides real-time monitoring and performance optimization of the mining operation boosted the mining operation's productivity and safety.","image":"home-product-4.webp","order_no":4}]`,
		IsJson:  true,
	}, {
		Section: "products",
		Key:     "is_active",
		Value:   "true",
		IsJson:  false,
	}, {
		Section: "impact",
		Key:     "desc-1",
		Value:   "Minerva's implementation leads to better utilization, higher output reduce carbon emission and improved safety.",
		IsJson:  false,
	}, {
		Section: "impact",
		Key:     "list-1",
		Value:   `[{"title":"9%","desc":"<span>Reducing <strong>mining operating costs</strong> through improved utilization, productivity, and mine plan compliance<span>","image":"home-impact-1-1.svg","order_no":1},{"title":"1.9-ton CO2e","desc":"<span>Additional prevent carbon emissions caused by transform traditional paperbased programs into a <strong>paperless SHE</strong> program and digital reports<span>","image":"home-impact-1-2.svg","order_no":2},{"title":"15%","desc":"<span><strong>Reduce carbon emission</strong> by reducing fuel ration, improving machine utilization, and productivity (decrease total of their fleet required to produce the output)<span>","image":"home-impact-1-3.svg","order_no":3},{"title":"21%","desc":"<span><strong>Decrease incident rates</strong> and improve overall safety performance by Realtime distribution of corrective actions and the generation of scorecards and analytics<span>","image":"home-impact-1-4.svg","order_no":4}]`,
		IsJson:  true,
	}, {
		Section: "impact",
		Key:     "desc-2",
		Value:   "We understand the key enablers of successful digital transformation and its prerequisites.",
		IsJson:  false,
	}, {
		Section: "impact",
		Key:     "list-2",
		Value:   `[{"title":"","desc":"To implement strong change management strategy to ensure successful digital technology adoption.","image":"home-impact-2-1.svg","order_no":1},{"title":"","desc":"To acquire suitable competencies with the necessary digital skills to operate, manage, and leverage the digital platform effectively.","image":"home-impact-2-2.svg","order_no":2},{"title":"","desc":"To reach a high level of maturity regarding mining operational excellence that provides a strong foundation in digital transformation implementation.","image":"home-impact-2-3.svg","order_no":3}]`,
		IsJson:  true,
	}, {
		Section: "impact",
		Key:     "is_active",
		Value:   "true",
		IsJson:  false,
	}, {
		Section: "case_studies",
		Key:     "id",
		Value:   "1",
		IsJson:  false,
	}, {
		Section: "case_studies",
		Key:     "is_active",
		Value:   "true",
		IsJson:  false,
	}, {
		Section: "testimonials",
		Key:     "list",
		Value:   `[{"content":"\"Aliquam erat volutpat. Donec in tortor justo. Nullam dapibus dolor vitae nibh aliquet lacinia. Phasellus vel sapien molestie, volutpat lorem in, iaculis diam. Nulla vel nunc magna. In hac habitasse platea dictumst. In mattis sollicitudin orci eget imperdiet. Nunc eget sagittis urna. Vestibulum eleifend quam elit, et fringilla magna dapibus eget.\"","name":"Nulla vucuaeasoo muka 1","position":"Suspendisse non ante tortor, mattis","photo":"home-testimonial-photo-1.webp","logo":"home-testimonial-logo-1.webp","order_no":1},{"content":"\"Aliquam erat volutpat. Donec in tortor justo. Nullam dapibus dolor vitae nibh aliquet lacinia. Phasellus vel sapien molestie, volutpat lorem in, iaculis diam. Nulla vel nunc magna. In hac habitasse platea dictumst. In mattis sollicitudin orci eget imperdiet. Nunc eget sagittis urna. Vestibulum eleifend quam elit, et fringilla magna dapibus eget.\"","name":"Nulla vucuaeasoo muka 2","position":"Suspendisse non ante tortor, mattis","photo":"home-testimonial-photo-2.webp","logo":"home-testimonial-logo-2.webp","order_no":2},{"content":"\"Aliquam erat volutpat. Donec in tortor justo. Nullam dapibus dolor vitae nibh aliquet lacinia. Phasellus vel sapien molestie, volutpat lorem in, iaculis diam. Nulla vel nunc magna. In hac habitasse platea dictumst. In mattis sollicitudin orci eget imperdiet. Nunc eget sagittis urna. Vestibulum eleifend quam elit, et fringilla magna dapibus eget.\"","name":"Nulla vucuaeasoo muka 3","position":"Suspendisse non ante tortor, mattis","photo":"home-testimonial-photo-3.webp","logo":"home-testimonial-logo-3.webp","order_no":3},{"content":"\"Aliquam erat volutpat. Donec in tortor justo. Nullam dapibus dolor vitae nibh aliquet lacinia. Phasellus vel sapien molestie, volutpat lorem in, iaculis diam. Nulla vel nunc magna. In hac habitasse platea dictumst. In mattis sollicitudin orci eget imperdiet. Nunc eget sagittis urna. Vestibulum eleifend quam elit, et fringilla magna dapibus eget.\"","name":"Nulla vucuaeasoo muka 4","position":"Suspendisse non ante tortor, mattis","photo":"home-testimonial-photo-4.webp","logo":"home-testimonial-logo-4.webp","order_no":4}]`,
		IsJson:  true,
	}, {
		Section: "testimonials",
		Key:     "is_active",
		Value:   "false",
		IsJson:  false,
	}}

	for i := 0; i < len(seeds); i++ {
		seeds[i].CreatedBy = helper.UserDummy.Email
		seeds[i].UpdatedBy = helper.UserDummy.Email
	}

	result := db.Create(&seeds)
	return result.Error
}
