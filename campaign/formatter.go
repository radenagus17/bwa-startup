package campaign

import "strings"

type CampaignFormatter struct {
	ID 								int `json:"id"`
	UserID 						int	`json:"user_id"`
	Name 							string `json:"name"`
	ShortDescription 	string `json:"short_description"`
	ImageURL 					string `json:"image_url"`
	GoalAmount 				int `json:"goal_amount"`
	CurrentAmount 		int `json:"current_amount"`
	Slug 							string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmoun
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter

}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter{
	
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns{
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID 								int 						`json:"id"`
	Name 							string 					`json:"name"`
	ShortDescription	string					`json:"short_description"`
	Description 			string					`json:"description"`
	ImageURL 					string					`json:"image_url"`
	GoalAmount 				int							`json:"goal_amount"`
	CurrentAmount 		int							`json:"current_amount"`
	UserID 						int 						`json:"user_id"`
	Slug							string					`json:"slug"`
	Perks							[]string				`json:"perks"`
	User 							CampaignUserFormatter `json:"user"`
	Images 						[]CampaignImageFormatter `json:"images"`
}


type CampaignUserFormatter struct {
	Name 						string `json:"name"`
	ImageURL 				string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL 			string `json:"image_url"`
	IsPrimary 		bool `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter{
	campaignDetailFormater := CampaignDetailFormatter{}
	campaignDetailFormater.ID = campaign.ID
	campaignDetailFormater.Name = campaign.Name
	campaignDetailFormater.ShortDescription = campaign.ShortDescription
	campaignDetailFormater.Description = campaign.Description
	campaignDetailFormater.ImageURL = ""
	campaignDetailFormater.GoalAmount = campaign.GoalAmoun
	campaignDetailFormater.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormater.Slug = campaign.Slug
	campaignDetailFormater.UserID = campaign.UserId

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormater.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _,perk := range strings.Split(campaign.Perks, ","){
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormater.Perks = perks

	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormater.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _,image := range campaign.CampaignImages{
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary=isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormater.Images = images

	return campaignDetailFormater
}

