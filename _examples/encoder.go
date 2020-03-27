// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/27 15:10:02

package main

import (
    "os"

    "github.com/FishGoddess/logit"
)

type myEncoder struct{}

// Customize you own encoder.
func (me *myEncoder) Encode(log *logit.Log) []byte {
    if log.Level() == logit.DebugLevel {
        return []byte("I am debug log ===> " + log.Msg() + "\n")
    }
    return []byte("Normal log +++> " + log.Msg() + "\n")
}

func main() {

    // logit.NewDefaultEncoder returns a default encoder with given time format.
    logger := logit.NewLogger(logit.DebugLevel, logit.NewDefaultHandler(os.Stdout, logit.NewDefaultEncoder("2006/01/02 15:04:05")))
    logger.Info("What time is it now?")

    // logit.NewJsonEncoder returns a json encoder with given time format and need formatting time.
    logger = logit.NewLoggerFrom(logit.Config{
        Level:   logit.DebugLevel,
        Writer:  os.Stdout,
        Encoder: logit.NewJsonEncoder("2006/01/02 15:04:05", false),
    })
    logger.Info("I am json log!")

    // You can customize you own encoder, see myEncoder.
    logger = logit.NewLoggerFrom(logit.Config{
        Level:   logit.DebugLevel,
        Writer:  os.Stdout,
        Encoder: &myEncoder{},
    })
    logger.Debug("Ha ha!")
    logger.Info("Hello!")
}