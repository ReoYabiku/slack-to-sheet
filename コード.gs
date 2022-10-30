function doPost(e) {
  const task = e.parameter.task;

  switch (task) {
    case "myFunc":
      myFunc();
      break;
    default:
      console.log("taskが違う");
  }
}

function myFunc() {
  const sheets = SpreadsheetApp.getActiveSpreadsheet();
  const sheet = sheets.getSheetByName('2022年10月');

  sheet.getRange(1, 3).setValue('hoge');
}
