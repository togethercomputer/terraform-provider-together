// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_translation

import (
	"bytes"
	"errors"
	"mime/multipart"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apiform"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type AudioTranslationModel struct {
	File                   types.String                                                `tfsdk:"file" json:"file,required"`
	Prompt                 types.String                                                `tfsdk:"prompt" json:"prompt,optional"`
	Language               types.String                                                `tfsdk:"language" json:"language,computed_optional"`
	Model                  types.String                                                `tfsdk:"model" json:"model,computed_optional"`
	ResponseFormat         types.String                                                `tfsdk:"response_format" json:"response_format,computed_optional"`
	Temperature            types.Float64                                               `tfsdk:"temperature" json:"temperature,computed_optional"`
	TimestampGranularities types.String                                                `tfsdk:"timestamp_granularities" json:"timestamp_granularities,computed_optional"`
	Duration               types.Float64                                               `tfsdk:"duration" json:"duration,computed"`
	Task                   types.String                                                `tfsdk:"task" json:"task,computed"`
	Text                   types.String                                                `tfsdk:"text" json:"text,computed"`
	Segments               customfield.NestedObjectList[AudioTranslationSegmentsModel] `tfsdk:"segments" json:"segments,computed"`
	Words                  customfield.NestedObjectList[AudioTranslationWordsModel]    `tfsdk:"words" json:"words,computed"`
}

func (r AudioTranslationModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		if e := writer.Close(); e != nil {
			err = errors.Join(err, e)
		}
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type AudioTranslationSegmentsModel struct {
	ID    types.Int64   `tfsdk:"id" json:"id,computed"`
	End   types.Float64 `tfsdk:"end" json:"end,computed"`
	Start types.Float64 `tfsdk:"start" json:"start,computed"`
	Text  types.String  `tfsdk:"text" json:"text,computed"`
}

type AudioTranslationWordsModel struct {
	End       types.Float64 `tfsdk:"end" json:"end,computed"`
	Start     types.Float64 `tfsdk:"start" json:"start,computed"`
	Word      types.String  `tfsdk:"word" json:"word,computed"`
	SpeakerID types.String  `tfsdk:"speaker_id" json:"speaker_id,computed"`
}
