// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_transcription

import (
	"bytes"
	"errors"
	"mime/multipart"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apiform"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type AudioTranscriptionModel struct {
	File                   types.String                                                         `tfsdk:"file" json:"file,required"`
	MaxSpeakers            types.Int64                                                          `tfsdk:"max_speakers" json:"max_speakers,optional"`
	MinSpeakers            types.Int64                                                          `tfsdk:"min_speakers" json:"min_speakers,optional"`
	Prompt                 types.String                                                         `tfsdk:"prompt" json:"prompt,optional"`
	Diarize                types.Bool                                                           `tfsdk:"diarize" json:"diarize,computed_optional"`
	Language               types.String                                                         `tfsdk:"language" json:"language,computed_optional"`
	Model                  types.String                                                         `tfsdk:"model" json:"model,computed_optional"`
	ResponseFormat         types.String                                                         `tfsdk:"response_format" json:"response_format,computed_optional"`
	Temperature            types.Float64                                                        `tfsdk:"temperature" json:"temperature,computed_optional"`
	TimestampGranularities types.String                                                         `tfsdk:"timestamp_granularities" json:"timestamp_granularities,computed_optional"`
	Duration               types.Float64                                                        `tfsdk:"duration" json:"duration,computed"`
	Task                   types.String                                                         `tfsdk:"task" json:"task,computed"`
	Text                   types.String                                                         `tfsdk:"text" json:"text,computed"`
	Segments               customfield.NestedObjectList[AudioTranscriptionSegmentsModel]        `tfsdk:"segments" json:"segments,computed"`
	SpeakerSegments        customfield.NestedObjectList[AudioTranscriptionSpeakerSegmentsModel] `tfsdk:"speaker_segments" json:"speaker_segments,computed"`
	Words                  customfield.NestedObjectList[AudioTranscriptionWordsModel]           `tfsdk:"words" json:"words,computed"`
}

func (r AudioTranscriptionModel) MarshalMultipart() (data []byte, contentType string, err error) {
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

type AudioTranscriptionSegmentsModel struct {
	ID    types.Int64   `tfsdk:"id" json:"id,computed"`
	End   types.Float64 `tfsdk:"end" json:"end,computed"`
	Start types.Float64 `tfsdk:"start" json:"start,computed"`
	Text  types.String  `tfsdk:"text" json:"text,computed"`
}

type AudioTranscriptionSpeakerSegmentsModel struct {
	ID        types.Int64                                                               `tfsdk:"id" json:"id,computed"`
	End       types.Float64                                                             `tfsdk:"end" json:"end,computed"`
	SpeakerID types.String                                                              `tfsdk:"speaker_id" json:"speaker_id,computed"`
	Start     types.Float64                                                             `tfsdk:"start" json:"start,computed"`
	Text      types.String                                                              `tfsdk:"text" json:"text,computed"`
	Words     customfield.NestedObjectList[AudioTranscriptionSpeakerSegmentsWordsModel] `tfsdk:"words" json:"words,computed"`
}

type AudioTranscriptionSpeakerSegmentsWordsModel struct {
	End       types.Float64 `tfsdk:"end" json:"end,computed"`
	Start     types.Float64 `tfsdk:"start" json:"start,computed"`
	Word      types.String  `tfsdk:"word" json:"word,computed"`
	SpeakerID types.String  `tfsdk:"speaker_id" json:"speaker_id,computed"`
}

type AudioTranscriptionWordsModel struct {
	End       types.Float64 `tfsdk:"end" json:"end,computed"`
	Start     types.Float64 `tfsdk:"start" json:"start,computed"`
	Word      types.String  `tfsdk:"word" json:"word,computed"`
	SpeakerID types.String  `tfsdk:"speaker_id" json:"speaker_id,computed"`
}
