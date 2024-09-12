require 'spec_helper'

RSpec.describe 'End To End Suite' do
  describe "full scenarios" do
    let(:commands) do
      [
          "create_warehouse_rack 6\n",
          "rack ZG11AQA 2024-02-28\n",
          "rack SD92349WW 2024-02-28\n",
          "rack ZG748WDG 2024-03-15\n",
          "rack KA887YHJ 2024-01-21\n",
          "rack KA888YHH 2024-04-01\n",
          "rack DG8789YH 2024-03-15\n",
          "rack_out 4\n",
          "status\n",
          "rack DL654ASA 2024-02-28\n",
          "rack DO123UJU 2024-02-28\n",
          "sku_numbers_for_product_with_exp_date 2024-02-28\n",
          "slot_numbers_for_product_with_exp_date 2024-02-28\n",
          "slot_number_for_sku_number DG8789YH\n",
          "slot_number_for_sku_number AB8875WF\n"
      ]
    end

    let(:expected) do
      [
          "Created a warehouse rack with 6 slots\n",
          "Allocated slot number: 1\n",
          "Allocated slot number: 2\n",
          "Allocated slot number: 3\n",
          "Allocated slot number: 4\n",
          "Allocated slot number: 5\n",
          "Allocated slot number: 6\n",
          "Slot number 4 is free\n",
          "Slot No.  SKU No.    Exp Date\n1         ZG11AQA   2024-02-28\n2         SD92349WW   2024-02-28\n3         ZG748WDG    2024-03-15\n5         KA888YHH    2024-04-01\n6         DG8789YH    2024-03-15\n",
          "Allocated slot number: 4\n",
          "Sorry, rack is full\n",
          "ZG11AQA, SD92349WW, DL654ASA\n",
          "1, 2, 4\n",
          "6\n",
          "Not found\n"
      ]
    end

    it "input from file" do
      pty = PTY.spawn("warehouse_rack #{File.join(File.dirname(__FILE__), '..', 'fixtures', 'file_input.txt')}")
      print 'Testing file input: '
      expect(fetch_stdout(pty)).to eq(expected.join(''))
    end

    it "interactive input" do
      pty = PTY.spawn("warehouse_rack")
      print 'Testing interactive input: '
      commands.each_with_index do |cmd, index|
        print cmd
        run_command(pty, cmd)
        expect(fetch_stdout(pty)).to end_with(expected[index])
      end
    end
  end
end
