require 'spec_helper'

RSpec.describe 'Warehouse Rack' do
  let(:pty) { PTY.spawn('warehouse_rack') }

  before(:each) do
    run_command(pty, "create_warehouse_rack 3\n")
  end

  it "can create a warehouse rack", :sample => true do
    expect(fetch_stdout(pty)).to end_with("Created a warehouse rack with 3 slots\n")
  end

  it "can rack a product" do
    run_command(pty, "rack ZG11AQA 2024-02-28\n")
    expect(fetch_stdout(pty)).to end_with("Allocated slot number: 1\n")
  end
  
  it "can rack out a product" do
    run_command(pty, "rack ZG11AQA 2024-02-28\n")
    run_command(pty, "rack_out 1\n")
    expect(fetch_stdout(pty)).to end_with("Slot number 1 is free\n")
  end
  
  it "can report status" do
    run_command(pty, "rack SD92349WW 2024-02-28\n")
    run_command(pty, "rack ZG11AQA 2024-02-28\n")
    run_command(pty, "rack ZG748WDG 2024-03-15\n")
    run_command(pty, "status\n")
    expect(fetch_stdout(pty)).to end_with(<<-EOTXT
Slot No.  SKU No.     Exp Date
1         ZG11AQA     2024-02-28
2         SD92349WW   2024-02-28
3         ZG748WDG    2024-03-15
EOTXT
)
  end
  
  pending "add more specs as needed"
end
