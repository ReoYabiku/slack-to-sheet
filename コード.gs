function doPost(e) {
  const task = e.parameter.task;

  switch (task) {
    case "myFunc":
      myFunc();
      break;
    case "select":
      const params = JSON.parse(e.postData.getDataAsString());
      select(row = params.row, column = params.column, value = params.value);
      break;
    default:
      console.log("taskが違う");
  }
}

function select(row, column, value) {
  const sheets = SpreadsheetApp.getActiveSpreadsheet();
  const sheet = sheets.getSheetByName('2022年10月');

  sheet.getRange(row, column).setValue(value);
}

function myFunc() {
  const sheets = SpreadsheetApp.getActiveSpreadsheet();
  const sheet = sheets.getSheetByName('2022年10月');

  sheet.getRange(1, 3).setValue('hoge');
}
