package entity

type Investor struct {
	ID             string
	Name           string
	AssertPosition []*InvestorAssertPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:             id,
		AssertPosition: []*InvestorAssertPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssertPosition) {
	i.AssertPosition = append(i.AssertPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) { //share: quota
	assertPosition := i.GetAssetPosition(assetID)
	if assertPosition == nil {
		i.AssertPosition = append(i.AssertPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assertPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetId string) *InvestorAssertPosition {
	for _, assertPosition := range i.AssertPosition {
		if assertPosition.AssertID == assetId {
			return assertPosition
		}
	}

	return nil
}

type InvestorAssertPosition struct {
	AssertID string
	Shares   int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssertPosition {
	return &InvestorAssertPosition{
		AssertID: assetID,
		Shares:   shares,
	}
}
