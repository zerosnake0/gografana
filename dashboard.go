package gografana

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type DashboardMeta struct {
	Type                  string    `json:"type,omitempty"`
	CanSave               bool      `json:"canSave,omitempty"`
	CanEdit               bool      `json:"canEdit,omitempty"`
	CanAdmin              bool      `json:"canAdmin,omitempty"`
	CanStar               bool      `json:"canStar,omitempty"`
	Slug                  string    `json:"slug,omitempty"`
	URL                   string    `json:"url,omitempty"`
	Expires               time.Time `json:"expires,omitempty"`
	Created               time.Time `json:"created,omitempty"`
	Updated               time.Time `json:"updated,omitempty"`
	UpdatedBy             string    `json:"updatedBy,omitempty"`
	CreatedBy             string    `json:"createdBy,omitempty"`
	Version               int       `json:"version,omitempty"`
	HasAcl                bool      `json:"hasAcl,omitempty"`
	IsFolder              bool      `json:"isFolder,omitempty"`
	FolderID              int       `json:"folderId,omitempty"`
	FolderTitle           string    `json:"folderTitle,omitempty"`
	FolderURL             string    `json:"folderUrl,omitempty"`
	Provisioned           bool      `json:"provisioned,omitempty"`
	ProvisionedExternalID string    `json:"provisionedExternalId,omitempty"`
}

type DashboardModelRequirement struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}

type DashboardAnnotation struct {
	BuiltIn    int    `json:"builtIn,omitempty"`
	DataSource string `json:"datasource,omitempty"`
	Enable     bool   `json:"enable,omitempty"`
	Hide       bool   `json:"hide,omitempty"`
	IconColor  string `json:"iconColor,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
}

type DashboardAnnotations struct {
	List []DashboardAnnotation `json:"list,omitempty"`
}

type DashboardLink struct {
	AsDropdown  bool     `json:"asDropdown,omitempty"`
	Icon        string   `json:"icon,omitempty"`
	IncludeVars bool     `json:"includeVars,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Title       string   `json:"title,omitempty"`
	Type        string   `json:"type,omitempty"`
}

type GridPos struct {
	H int `json:"h,omitempty"`
	W int `json:"w,omitempty"`
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

type Legend struct {
	AlignAsTable bool   `json:"alignAsTable,omitempty"`
	Avg          bool   `json:"avg,omitempty"`
	Current      bool   `json:"current,omitempty"`
	HideEmpty    bool   `json:"hideEmpty,omitempty"`
	HideZero     bool   `json:"hideZero,omitempty"`
	Max          bool   `json:"max,omitempty"`
	Min          bool   `json:"min,omitempty"`
	RightSide    bool   `json:"rightSide,omitempty"`
	Show         bool   `json:"show,omitempty"`
	Sort         string `json:"sort,omitempty"`
	SortDesc     bool   `json:"sortDesc,omitempty"`
	Total        bool   `json:"total,omitempty"`
	Values       bool   `json:"values,omitempty"`
}

type PanelTarget struct {
	Expr           string `json:"expr,omitempty"`
	Format         string `json:"format,omitempty"`
	Interval       string `json:"interval,omitempty"`
	IntervalFactor int    `json:"intervalFactor,omitempty"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	RefID          string `json:"refId,omitempty"`
	Step           int    `json:"step,omitempty"`
}

type ToolTip struct {
	MsResolution bool   `json:"msResolution,omitempty"`
	Shared       bool   `json:"shared,omitempty"`
	Sort         int    `json:"sort,omitempty"`
	ValueType    string `json:"value_type,omitempty"`
}

type XAxis struct {
	Mode string `json:"mode,omitempty"`
	Show bool   `json:"show,omitempty"`

	// TODO:
	Buckets json.RawMessage `json:"buckets,omitempty"`
	Name    json.RawMessage `json:"name,omitempty"`
	Values  json.RawMessage `json:"values,omitempty"`
}

type YAxe struct {
	Format  string `json:"format,omitempty"`
	LogBase int    `json:"logBase,omitempty"`
	Show    bool   `json:"show,omitempty"`

	// TODO:
	Label json.RawMessage `json:"label,omitempty"`
	Max   json.RawMessage `json:"max,omitempty"`
	Min   json.RawMessage `json:"min,omitempty"`
}

type DashboardPanel struct {
	Collapsed     bool             `json:"collapsed,omitempty"`
	GridPos       GridPos          `json:"gridPos,omitempty"`
	ID            int              `json:"id,omitempty"`
	Panels        []DashboardPanel `json:"panels,omitempty"`
	Title         string           `json:"title,omitempty"`
	Transparent   bool             `json:"transparent,omitempty"`
	Type          string           `json:"type,omitempty"`
	Bars          bool             `json:"bars,omitempty"`
	DashLength    int              `json:"dashLength,omitempty"`
	Dashes        bool             `json:"dashes,omitempty"`
	DataSource    string           `json:"datasource,omitempty"`
	Description   string           `json:"description,omitempty"`
	Editable      bool             `json:"editable,omitempty"`
	Error         bool             `json:"error,omitempty"`
	Fill          int              `json:"fill,omitempty"`
	Legend        Legend           `json:"legend,omitempty"`
	Lines         bool             `json:"lines,omitempty"`
	LineWidth     int              `json:"linewidth,omitempty"`
	NullPointMode string           `json:"nullPointMode,omitempty"`
	Percentage    bool             `json:"percentage,omitempty"`
	PointRadius   int              `json:"pointradius,omitempty"`
	Points        bool             `json:"points,omitempty"`
	Renderer      string           `json:"renderer,omitempty"`
	SpaceLength   int              `json:"spaceLength,omitempty"`
	Stack         bool             `json:"stack,omitempty"`
	SteppedLine   bool             `json:"steppedLine,omitempty"`
	Targets       []PanelTarget    `json:"targets,omitempty"`
	ToolTip       ToolTip          `json:"tooltip,omitempty"`
	XAxis         XAxis            `json:"xaxis,omitempty"`
	YAxes         []YAxe           `json:"yaxes,omitempty"`

	// TODO:
	Alerting        json.RawMessage `json:"alerting,omitempty"`
	Repeat          json.RawMessage `json:"repeat,omitempty"`
	Links           json.RawMessage `json:"links,omitempty"`
	AliasColors     json.RawMessage `json:"aliasColors,omitempty"`
	Grid            json.RawMessage `json:"grid,omitempty"`
	SeriesOverrides json.RawMessage `json:"seriesOverrides,omitempty"`
	Thresholds      json.RawMessage `json:"thresholds,omitempty"`
	TimeFrom        json.RawMessage `json:"timeFrom,omitempty"`
	TimeShift       json.RawMessage `json:"timeShift,omitempty"`
}

type DashboardTemplatingItem struct {
	Hide           int      `json:"hide,omitempty"`
	Name           string   `json:"name,omitempty"`
	Query          string   `json:"query,omitempty"`
	Refresh        int      `json:"refresh,omitempty"`
	Regex          string   `json:"regex,omitempty"`
	Type           string   `json:"type,omitempty"`
	IncludeAll     bool     `json:"includeAll,omitempty"`
	Label          string   `json:"label,omitempty"`
	Multi          bool     `json:"multi,omitempty"`
	Sort           int      `json:"sort,omitempty"`
	TagValuesQuery string   `json:"tagValuesQuery,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	TagsQuery      string   `json:"tagsQuery,omitempty"`
	UseTags        bool     `json:"useTags,omitempty"`

	// TODO:
	Options    json.RawMessage `json:"options,omitempty"`
	AllValue   json.RawMessage `json:"allValue,omitempty"`
	DataSource string          `json:"datasource,omitempty"`
	Current    json.RawMessage `json:"current,omitempty"`
}

type DashboardTemplating struct {
	List []DashboardTemplatingItem `json:"list,omitempty"`
}

type DashboardTime struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type TimePicker struct {
	RefreshIntervals []string `json:"refresh_intervals,omitempty"`
	TimeOptions      []string `json:"time_options,omitempty"`
}

type DashboardModel struct {
	Requires      []DashboardModelRequirement `json:"__requires,omitempty"`
	Annotations   DashboardAnnotations        `json:"annotations,omitempty"`
	Description   string                      `json:"description,omitempty"`
	Editable      bool                        `json:"editable,omitempty"`
	GNetID        int                         `json:"gnetId,omitempty"`
	GraphToolTip  int                         `json:"graphTooltip,omitempty"`
	ID            int                         `json:"id,omitempty"`
	Iteration     int                         `json:"iteration,omitempty"`
	Links         []DashboardLink             `json:"links,omitempty"`
	Panels        []DashboardPanel            `json:"panels,omitempty"`
	SchemaVersion int                         `json:"schemaVersion,omitempty"`
	Style         string                      `json:"style,omitempty"`
	Tags          []string                    `json:"tags,omitempty"`
	Templating    DashboardTemplating         `json:"templating,omitempty"`
	Time          DashboardTime               `json:"time,omitempty"` // required when creating dashboard
	TimePicker    TimePicker                  `json:"timepicker,omitempty"`
	Timezone      string                      `json:"timezone,omitempty"`
	Title         string                      `json:"title,omitempty"`
	UID           string                      `json:"uid,omitempty"`
	Version       int                         `json:"version,omitempty"`
}

type Dashboard struct {
	Meta DashboardMeta `json:"meta,omitempty"`

	Dashboard DashboardModel `json:"dashboard,omitempty"`
}

func (c *Client) GetDashboardByUID(ctx context.Context, uid string) (db Dashboard, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/dashboards/uid/"+uid, nil)
	if err != nil {
		return
	}
	err = c.doJsonRequest200(req, &db)
	return
}

type DashboardCreationRequest struct {
	Dashboard DashboardModel `json:"dashboard,omitempty"`

	FolderID int `json:"folderId,omitempty"`

	Overwrite bool `json:"overwrite,omitempty"`
}

type DashboardCreationResponse struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Status  string `json:"status"`
	UID     string `json:"uid"`
	URL     string `json:"url"`
	Version int    `json:"version"`
}

func (c *Client) CreateDashboard(ctx context.Context, dcReq *DashboardCreationRequest) (res DashboardCreationResponse,
	err error) {
	b, err := jsonMarshal(dcReq)
	if err != nil {
		return
	}
	req, err := c.newRequest(ctx, http.MethodPost, "/api/dashboards/db", bytes.NewReader(b))
	if err != nil {
		return
	}
	setRequestContentTypeJson(req)
	err = c.doJsonRequest200(req, &res)
	return
}
