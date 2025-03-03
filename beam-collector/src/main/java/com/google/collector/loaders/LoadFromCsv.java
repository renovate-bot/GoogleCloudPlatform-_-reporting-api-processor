/*
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.google.collector.loaders;

import static reports.SecurityReportOuterClass.*;

import com.google.common.collect.ImmutableList;
import java.util.Scanner;
import org.apache.beam.sdk.transforms.DoFn;
import reports.CspReportOuterClass.CspReport;
import reports.SecurityReportOuterClass.SecurityReport.Disposition;

/**
 * This transform can load reports from CSV files used during debugging. The CSV representation of
 * reports is discouraged in production.
 */
final class LoadFromCsv extends DoFn<String, SecurityReport> {

  @ProcessElement
  public void processElement(ProcessContext ctx) throws IllegalArgumentException {
    ImmutableList<String> columns = loadColumns(ctx);
    // hacky skip of unwanted lines: header, empties, unparsed
    if (columns.isEmpty() || columns.size() <= 12 || columns.get(1).equals("report_time")) {
      return;
    }

    CspReport cspReport = CspReport.newBuilder()
        .setReferrer(columns.get(4))
        .setBlockedUri(columns.get(5))
        .setViolatedDirective(columns.get(6))
        .setOriginalPolicy(columns.get(7))
        .setSourceFile(columns.get(8))
        .setEffectiveDirective(columns.get(13))
        .setScriptSample(columns.get(19))
        .setLineNumber(-1)
        .setColumnNumber(-1)
        .build();

    SecurityReport securityReport = SecurityReport.newBuilder()
        .setReportChecksum("foo")
        .setReportTime(1L)
        .setReportCount(1L)
        .setUserAgent("foo")
        .setDisposition(Disposition.REPORTING)
        .setCspReport(cspReport)
        .build();
    ctx.output(securityReport);
  }

  private ImmutableList<String> loadColumns(ProcessContext ctx) {
    ImmutableList.Builder<String> values = ImmutableList.builder();
    try (Scanner rowScanner = new Scanner(
        ctx.element().replaceAll("\n", "")
    )) {
      rowScanner.useDelimiter(",");
      while (rowScanner.hasNext()) {
        values.add(rowScanner.next());
      }
    }

    return values.build();
  }
}