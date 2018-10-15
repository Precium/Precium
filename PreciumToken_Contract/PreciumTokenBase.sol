pragma solidity ^0.4.24;

import "./IERC20.sol";
import "./SafeMath.sol";

contract PreciumTokenBase is IERC20 {
    using SafeMath for uint256;

    mapping (address => uint256) public balances_;
    mapping (address => mapping (address => uint256)) public allowed_;

    uint256 public totalSupply_;

    function totalSupply() public view returns (uint256) {
        return totalSupply_;
    }

    function balanceOf(address _owner) public view returns (uint256) {
        return balances_[_owner];
    }

    function allowance(address _owner, address _spender) public view returns (uint256) {
        return allowed_[_owner][_spender];
    }

    function _transfer(address _from, address _to, uint256 _value) internal returns (bool) {
        require(_value <= balances_[_from]);
        require(_to != address(0));

        balances_[_from] = balances_[_from].sub(_value);
        balances_[_to] = balances_[_to].add(_value);
        emit Transfer(_from, _to, _value);
        
        return true;
    }

    function transfer(address _to, uint256 _value) public returns (bool) {
        return _transfer(msg.sender, _to, _value);
    }

    function approve(address _spender, uint256 _value) public returns (bool) {
        allowed_[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }

    function _transferFrom(address _from, address _to, uint256 _value) internal returns (bool) {
        require(_value <= balances_[_from]);
        require(_value <= allowed_[_from][msg.sender]);
        require(_to != address(0));

        balances_[_from] = balances_[_from].sub(_value);
        balances_[_to] = balances_[_to].add(_value);
        allowed_[_from][msg.sender] = allowed_[_from][msg.sender].sub(_value);
        emit Transfer(_from, _to, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        return _transferFrom(_from, _to, _value);
    }

    function increaseApproval(address _spender, uint256 _addedValue) public returns (bool) {
        allowed_[msg.sender][_spender] = allowed_[msg.sender][_spender].add(_addedValue);
        emit Approval(msg.sender, _spender, allowed_[msg.sender][_spender]);
        return true;
    }

    function decreaseApproval(address _spender, uint256 _subtractedValue) public returns (bool) {
        uint256 oldValue = allowed_[msg.sender][_spender];
        if (_subtractedValue >= oldValue) {
            allowed_[msg.sender][_spender] = 0;
        } else {
            allowed_[msg.sender][_spender] = oldValue.sub(_subtractedValue);
        }
        emit Approval(msg.sender, _spender, allowed_[msg.sender][_spender]);
        return true;
    }

    function _burn(address _account, uint256 _amount) internal {
        require(_account != 0);
        require(_amount <= balances_[_account]);

        totalSupply_ = totalSupply_.sub(_amount);
        balances_[_account] = balances_[_account].sub(_amount);
        emit Transfer(_account, address(0), _amount);
    }
    
    function burn(uint256 _amount) public {
        _burn(msg.sender, _amount);
    }
}
