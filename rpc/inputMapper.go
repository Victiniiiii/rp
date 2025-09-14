package rpc

func mapActivity(activity *Activity) *PayloadActivity {
	final := &PayloadActivity{
		Details: activity.Details,
		State:   activity.State,
		Assets: PayloadAssets{
			LargeImage: activity.LargeImage,
			LargeText:  activity.LargeText,
			SmallImage: activity.SmallImage,
			SmallText:  activity.SmallText,
		},
		Type: 2,
	}

	if activity.Timestamps != nil && activity.Timestamps.Start != nil {
		start := uint64(activity.Timestamps.Start.UnixNano() / 1e6)
		final.Timestamps = &PayloadTimestamps{
			Start: &start,
		}
		if activity.Timestamps.End != nil {
			end := uint64(activity.Timestamps.End.UnixNano() / 1e6)
			final.Timestamps.End = &end
		}
	}

	if activity.Party != nil {
		final.Party = &PayloadParty{
			ID:   activity.Party.ID,
			Size: [2]int{activity.Party.Players, activity.Party.MaxPlayers},
		}
	}

	if activity.Secrets != nil {
		final.Secrets = &PayloadSecrets{
			Join:     activity.Secrets.Join,
			Match:    activity.Secrets.Match,
			Spectate: activity.Secrets.Spectate,
		}
	}

	if len(activity.Buttons) > 0 {
		for _, btn := range activity.Buttons {
			final.Buttons = append(final.Buttons, &PayloadButton{
				Label: btn.Label,
				Url:   btn.Url,
			})
		}
	}

	return final
}
